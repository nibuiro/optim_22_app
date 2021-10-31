package main

import (
  "flag"
  "os"
  "fmt"
  "time"
  "net/http"
  "context"
  "os/signal"
  "syscall"
  "gorm.io/gorm"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/pprof"
  "github.com/gin-contrib/zap"
  "github.com/gin-contrib/cors"
  "optim_22_app/model"
  "optim_22_app/typefile"
  "optim_22_app/pkg/log"
  "optim_22_app/internal/pkg/auth22"
  "optim_22_app/internal/pkg/config"
  "optim_22_app/internal/app/home"
  "optim_22_app/internal/app/client"
  "optim_22_app/internal/app/request"
  "optim_22_app/internal/app/submission"
  "optim_22_app/internal/app/engineer"
  "optim_22_app/internal/app/user"
  "optim_22_app/internal/app/profile"  
  "optim_22_app/internal/app/comment"  
)

//このファイルmain.goの引数の定義
var flagConfig = flag.String("config", "./configs/app.yaml", "Appの設定ファイルへのパス")

func main() {
  //Ctrl+CなどのOSからの割り込み信号を待機させるリスナ, それに対する応答を定義
  ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
  //main終了時 //リスナを削除 //NotifyContextへのバイパスを解除
  defer stop()

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

  // DB接続後、マイグレーションを実行する。
  // 手順としては、まずコンテナを立ち上げた後、mysqlでoptim_devデータベースを作成する。
  // その後、model.InitDB(),import(optim_22_app/model)のコメントを外し、カレントディレクトリでgo run main.goを実行する。
  // プログラムの詳細はmodel/migrate.goに記載。
  model.InitDB()

  // マイグレーションは定義したstructをAutoMigrateの引数に渡すことで、
  // それに対応するテーブルの作成を行う。
  // テーブル作成時にオプションを付けたい場合、db.Set()を利用する。
  model.Db.AutoMigrate(
    &typefile.User{},
    &typefile.Profile{},
    &typefile.Client{},
    &typefile.Engineer{},
    &typefile.Winner{},
    &typefile.Request{},
    &typefile.Comment{},
    &typefile.Submission{},
  )
  // テスト実行前に利用するデータを作成する
  model.CreateTestData()


  //#region HTTPサーバをビルド
  address := fmt.Sprintf(":%v", cfg.ServerPort)

  srv := &http.Server{
    Addr:              address,
    Handler:           buildHandler(model.Db, logger, cfg),
    ReadTimeout:       time.Duration(cfg.ReadTimeout * int64(time.Second)),
    ReadHeaderTimeout: time.Duration(cfg.ReadHeaderTimeout * int64(time.Second)),
    WriteTimeout:      time.Duration(cfg.WriteTimeout * int64(time.Second)),
    IdleTimeout:       time.Duration(cfg.IdleTimeout * int64(time.Second)),
    MaxHeaderBytes:    1<<20,
  }
  //#endregion 

  //httpサーバスレッドを発行
  go func() {
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      logger.Errorf("listen: %s\n", err)
    }
  }()
  //割り込みシグナルを待機
  <-ctx.Done()

  //リスナを削除
  stop()

  //タイムアウトつきコンテキストを発行
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  //新たなHTTPセッションを受け付けないようにロックする
  if err := srv.Shutdown(ctx); err != nil {
    logger.Info("Server forced to shutdown: ", err)
  }

  logger.Info("Server exiting")
}



//任意のポートについてのHTTPハンドラを構築
func buildHandler(db *gorm.DB, logger log.Logger, cfg *config.Config) http.Handler {
  //[GIN-DEBUG]出力を無効化
  gin.SetMode(gin.ReleaseMode)
  //ミドルウェアが接続されていない新しい空のEngineインスタンスを取得 //担当：石森
  //!! Default()は、LoggerとRecoveryのミドルウェアが既にアタッチされているEngineインスタンスを返す
  e := gin.New()
  e.Use(CORS())
  //ginのログをloggerでとる //フォーマット形式はloggerに依存する //担当：石森
  e.Use(ginzap.Ginzap(logger.Desugar(), time.RFC3339, true))
  //パニック時ステータスコード500を送出 //担当：石森
  e.Use(ginzap.RecoveryWithZap(logger.Desugar(), true))
  pprof.Register(e)

  //#region 認証機能群
  authRepository := auth22.NewRepository(db, logger)
  authService := auth22.NewService(cfg, authRepository, logger)
  auth := auth22.New(authService, "localhost")
  //アクセストークンとリフレッシュトークンの発行
  e.POST("/auth", auth.Login())
  //トークンのリフレッシュ
  e.POST("/auth/refresh_token", auth.RefreshAccessTokenAndRefreshToken())
  //許可されたメソッドとパスのペア以外についてアクセストークンを検証
  e.Use(auth.ValidateAccessToken(auth22.GetRule(), true))
  //#endregion //担当：石森

  // homepageを表示するハンドラ
  e.GET("/api/requests",home.ShowHomepage)
  // NewRequestで得たengineer_idとrequest_idによって、エンジニアが特定リクエストに参加することをデータベースに登録するためのハンドラ
  e.POST("/api/request",client.CreateRequest)
  // request_idをparamにして特定リクエストの詳細を表示する。
  e.GET("/api/request/:request_id",request.ShowRequest)
  // クライアントが編集したリクエストを更新できるようにするハンドラ
  e.PUT("/api/request/:request_id",client.UpdateRequest)
  // JoinRequestで得たデータによって、エンジニアが特定リクエストに参加することをデータベースに登録するためのハンドラ
  e.POST("/api/request/:request_id",engineer.CreateEngineerJoin)

  // 特定リクエストのサブミッション一覧ページから勝者を選択できるようにするハンドラ
  e.POST("/api/winner/:request_id",client.DecideWinner)

  // submission_idをparamにして特定サブミッションの詳細を表示する。
  e.GET("/api/submission/:submission_id",submission.ShowSubmission)
  // エンジニアが編集したsubmissionを更新できるようにするハンドラ
  e.PUT("/api/submission/:submission_id",engineer.UpdateSubmission)
  // NewSubmissionで得たデータによって、エンジニアがsubmissionを提出したことをデータベースに登録するためのハンドラ
  e.POST("/api/submission/:request_id",engineer.CreateSubmission)

  
  //#region ユーザエンドポイントの構築
  userRepository := user.NewRepository(db, logger)
  userService := user.NewService(userRepository, logger)
  user.RegisterHandlers(e.Group("/api/user"), userService, logger)
  //#endregion //担当：石森
  
  //#region プロフィールエンドポイントの構築
  profileRepository := profile.NewRepository(db, logger)
  profileService := profile.NewService(profileRepository, logger)
  profile.RegisterHandlers(e.Group("/api/user"), profileService, logger)
  //#endregion //担当：石森

  //#region ディスカッションエンドポイントの構築
  commentRepository := comment.NewRepository(db, logger)
  commentService := comment.NewService(commentRepository, logger)
  comment.RegisterHandlers(e.Group("/api/discussion/"), commentService, logger)
  //#endregion //担当：石森

  return e
}

//担当：石森
func CORS() gin.HandlerFunc {
  return cors.New(cors.Config{
    // アクセスを許可したいHTTPメソッド
    AllowMethods: []string{
      "*",
    },
    // 許可したいHTTPリクエストヘッダ
    AllowHeaders: []string{
      "Access-Control-Allow-Credentials",
      "Access-Control-Allow-Headers",
      "Content-Type",
      "Content-Length",
      "Accept-Encoding",
      "Authorization",
      "Refresh-Token",
    },
    ExposeHeaders: []string{
      "Authorization",
      "Refresh-Token",
    },
    AllowOrigins: []string{
      "*",
    },
    // cookieなどの情報を必要とするかどうか
    AllowCredentials: true, //2021/10/08-21:29現時点では必要ない
    // preflightリクエストの結果をキャッシュする時間
    MaxAge: 24 * time.Hour,
  })
}
