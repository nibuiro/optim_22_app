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
  "optim_22_app/internal/hello"
  "optim_22_app/internal/client"
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
  e.LoadHTMLGlob("views/*.html")

  // ハンドラの指定
  e.GET("/hello", hello.Hello)

  // ハンドラの指定
  e.GET("/newhello", hello.NewHello)

  client_e := e.Group("/client")
  {
    client_e.GET("/new_request", client.NewRequest)
    client_e.POST("/create_request",client.CreateRequest)
    // client_idはサーバーサイドで直接取得できると捉えているため、開発後はクエリパラメータに入れない。
    client_e.GET("/show_request/:client_id", client.ShowRequest)
    // request_idをparamにして特定リクエストのサブミッションを表示するハンドラ
    client_e.GET("/show_submission/:request_id",client.ShowSubmission)
    // 特定リクエストのサブミッション一覧ページから勝者を選択できるようにするハンドラ
    client_e.POST("/decide_winner",client.DecideWinner)
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