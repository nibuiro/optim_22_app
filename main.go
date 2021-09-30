package main

import (
  "flag"
  "os"
  "fmt"
  "time"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/zap"
  "golang.org/x/sync/errgroup"
  "optim_22_app/model"
  "optim_22_app/typefile"
  "optim_22_app/pkg/log"
  "optim_22_app/internal/pkg/config"
  "optim_22_app/internal/home"
  "optim_22_app/internal/client"
  "optim_22_app/internal/request"
  "optim_22_app/internal/submission"
  "optim_22_app/internal/engineer"
)

var (
    g errgroup.Group
)

//このファイルmain.goの引数の定義
var flagConfig = flag.String("config", "./configs/app.yaml", "Appの設定ファイルへのパス")

func main() {
  //引数をパース
  flag.Parse()

  //zapロガーを設定ファイル`config/zap.yaml`を元に取得
  logger := log.New()
  logger.Debugf("start app")

  // load application configurations
  cfg, err := config.Load(*flagConfig, logger)
  if err != nil {
    logger.Errorf("failed to load application configuration: %s", err)
    os.Exit(-1)
  }


  logger.Debugf(cfg.DSN)
  

  // DB接続後、マイグレーションを実行する。
  // 手順としては、まずコンテナを立ち上げた後、mysqlでoptim_devデータベースを作成する。
  // その後、model.InitDB(),import(optim_22_app/model)のコメントを外し、カレントディレクトリでgo run main.goを実行する。
  // プログラムの詳細はmodel/migrate.goに記載。
  model.InitDB()

  // マイグレーションは定義したstructをAutoMigrateの引数に渡すことで、
  // それに対応するテーブルの作成を行う。
  // テーブル作成時にオプションを付けたい場合、db.Set()を利用する。
  model.Db.AutoMigrate(&typefile.User{},&typefile.Client{},&typefile.Engineer{},&typefile.Winner{},&typefile.Request{},&typefile.Submission{})
  // テスト実行前に利用するデータを作成する
  model.CreateTestData()


  //#region HTTPサーバをビルド
  address := fmt.Sprintf(":%v", cfg.ServerPort)

  hs := &http.Server{
    Addr:    address,
    Handler: buildHandler(logger, cfg), //, dbcontext.New(db)
  }
  //#endregion

  g.Go(func() error {
      return hs.ListenAndServe()
  })

  if err := g.Wait(); err != nil {
      logger.Error(err)
  }

}


//任意のポートについてのHTTPハンドラを構築
func buildHandler(logger log.Logger, cfg *config.Config) http.Handler { //, db *dbcontext.DB

  //ミドルウェアが接続されていない新しい空のEngineインスタンスを取得
  //!! Default()は、LoggerとRecoveryのミドルウェアが既にアタッチされているEngineインスタンスを返す
  e := gin.New()
  //ginのログをloggerでとる //フォーマット形式はloggerに依存する
  e.Use(ginzap.Ginzap(logger.Desugar(), time.RFC3339, true))
  //パニック時ステータスコード500を送出
  e.Use(ginzap.RecoveryWithZap(logger.Desugar(), true))

  // 事前にテンプレートをロード
  e.LoadHTMLGlob("views/*/*.html")

  // homepageを標示するハンドラ
  e.GET("/", home.ShowHomepage)

  client_e := e.Group("/client")
  {
    // NewRequestで得たengineer_idとrequest_idによって、エンジニアが特定リクエストに参加することをデータベースに登録するためのハンドラ
    client_e.POST("/create_request",client.CreateRequest)
    // client_idはサーバーサイドで直接取得できると捉えているため、開発後はクエリパラメータに入れない。
    client_e.GET("/show_request/:client_id", client.ShowRequest)
    // request_idをparamにして特定リクエストのサブミッションを表示するハンドラ
    client_e.GET("/show_submission/:request_id",client.ShowSubmission)
    // 特定リクエストのサブミッション一覧ページから勝者を選択できるようにするハンドラ
    client_e.POST("/decide_winner",client.DecideWinner)
    // クライアントが編集したリクエストを更新できるようにするハンドラ
    client_e.POST("/update_request",client.UpdateRequest)
  }

  request_e := e.Group("/request")
  {
    // request_idをparamにして特定リクエストの詳細を表示する。
    request_e.GET("/show_request/:request_id",request.ShowRequest)
  }

  submission_e := e.Group("/submission")
  {
    // request_idをparamにして特定リクエストの詳細を表示する。
    submission_e.GET("/show_submission/:submission_id",submission.ShowSubmission)
  }

  engineer_e := e.Group("/engineer")
  {
    // JoinRequestで得たデータによって、エンジニアが特定リクエストに参加することをデータベースに登録するためのハンドラ
    engineer_e.POST("/create_engineer_join",engineer.CreateEngineerJoin)
    // NewSubmissionで得たデータによって、エンジニアがsubmissionを提出したことをデータベースに登録するためのハンドラ
    engineer_e.POST("/create_submission",engineer.CreateSubmission)
    // engineer_idはサーバーサイドで直接取得できると捉えているため、開発後はクエリパラメータに入れない。
    engineer_e.GET("/show_join_request/:engineer_id", engineer.ShowJoinRequest)
    // エンジニアが編集したsubmissionを更新できるようにするハンドラ
    engineer_e.POST("/update_submission",engineer.UpdateSubmission)
  }

  e.NoRoute(func(c *gin.Context) {
    c.HTML(http.StatusOK, "error404.html", gin.H{})})

  //authHandler := auth.Handler(cfg.JWTSigningKey)

//  user.RegisterHandlers(rg.Group(""),
//    user.NewService(user.NewRepository(db, logger), logger),
//    authHandler, logger,
//  )
//
//  auth.RegisterHandlers(rg.Group(""),
//    auth.NewService(cfg.JWTSigningKey, cfg.JWTExpiration, logger),
//    logger,
//  )

  return e
}