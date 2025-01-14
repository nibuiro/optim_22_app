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

/*
 * [Go Gin Graceful-Shutdownについて](https://sourjp.github.io/posts/go-gin-graceful/)
 * [examples/graceful-shutdown at master · gin-gonic/examples · GitHub](https://github.com/gin-gonic/examples/tree/master/graceful-shutdown)
 * これらを参考にGraceful-Shutdownを実装
 */

func main() {
  //Ctrl+CなどのOSからの割り込み信号を待機させるリスナ, それに対する応答を定義
  ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
  //main終了時 //リスナを削除 //NotifyContextへのバイパスを解除
  /*
   * * [signal package - os/signal - pkg.go.dev](https://pkg.go.dev/os/signal#NotifyContext)
   *   > The stop function unregisters the signal behavior, 
   *   > which, like signal.Reset, may restore the default behavior
   *   > for a given signal. For example, the default behavior of a 
   *   > Go program receiving os.Interrupt is to exit. Calling 
   *   > NotifyContext(parent, os.Interrupt) will change the behavior 
   *   > to cancel the returned context. Future interrupts received 
   *   > will not trigger the default (exit) behavior until the returned 
   *   > stop function is called.
   *
   * * [Go1.7のcontextパッケージ | Taichi Nakashima](https://deeeet.com/writing/2016/07/22/context/)
   * contextについての参考文献
   */
  defer stop()

  //引数をパース
  flag.Parse()

  /* 
   * * [uber-go/zap： Blazing fast, structured, leveled logging in Go.](https://github.com/uber-go/zap)
   *   Why zap ? 
   *   zapがユースケースの多さ、開発元がuberであるなどの理由
   *   速く開発するためエラー出力について悩む必要がないようにSugaredLoggerを採用
   */

  //zapロガーを設定ファイル`config/zap.yaml`を元に取得
  logger := log.New()
  logger.Debugf("start app")

  //設定をロード
  cfg, err := config.Load(*flagConfig, logger)
  if err != nil {
    logger.Errorf("failed to load application configuration: %s", err)
    //未知のエラー
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

  /*
   * * [http.ListenAndServe() をインターネットに公開してはいけない - Qiita](https://qiita.com/methane/items/2453ed86305f6950775b)
   *   タイムアウトを適切に設定しなければリソースリークの危険性がある。
  */

  srv := &http.Server{
    Addr:              address,
    Handler:           buildHandler(model.Db, logger, cfg),
    ReadTimeout:       time.Duration(cfg.ReadTimeout * int64(time.Second)),
    ReadHeaderTimeout: time.Duration(cfg.ReadHeaderTimeout * int64(time.Second)),
    WriteTimeout:      time.Duration(cfg.WriteTimeout * int64(time.Second)),
    IdleTimeout:       time.Duration(cfg.IdleTimeout * int64(time.Second)),
    //[http package - net/http - pkg.go.dev](https://pkg.go.dev/net/http)
    //本来であれば考えうる最大サイズを計算して設定すべき
    MaxHeaderBytes:    1<<20,
  }
  //#endregion 

  //httpサーバスレッドを発行 //ショットダウン・クローズ処理がされていなければエラーを出力
  /*
   * * [http package - net/http - pkg.go.dev](https://pkg.go.dev/net/http#ErrServerClosed)
   *   >  ErrServerClosed is returned by the Server's Serve, ServeTLS, 
   *   >  ListenAndServe, and ListenAndServeTLS methods after a call to 
   *   >  Shutdown or Close.
   */
  /*
   * * [Channel_types](https://golang.org/ref/spec#Channel_types)
   * * [Go by Example： Channels](https://gobyexample.com/channels)
   * * [Go by Example： Select](https://gobyexample.com/select)
   * チャンネルについての参考
  */
  go func() {
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      logger.Errorf("listen: %s\n", err)
    }
  }()
  //割り込みシグナルを待機 
  <-ctx.Done()

  //リスナを削除
  stop()

  /*
   * * [context package - context - pkg.go.dev](https://pkg.go.dev/context#Background)
   *   > Background returns a non-nil, empty Context. 
   *   > It is never canceled, has no values, and has no deadline. 
   *   > It is typically used by the main function, initialization, 
   *   > and tests, and as the top-level Context for incoming requests.
   * タイマ付きコンテキストを注入しタイムアウトより前に処理が終わらなかった場合
   * httpサーバスレッドを強制停止
   */
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
  //[GIN-DEBUG]出力を無効化 //stdoutは高負荷
  gin.SetMode(gin.ReleaseMode)
  //ミドルウェアが接続されていない新しい空のEngineインスタンスを取得 //担当：石森
  //!! Default()は、LoggerとRecoveryのミドルウェアが既にアタッチされているEngineインスタンスを返す
  e := gin.New()
  /*
   * * [gin package - github.com/gin-gonic/gin - pkg.go.dev](https://pkg.go.dev/github.com/gin-gonic/gin#Default)
   * * [gin package - github.com/gin-gonic/gin - pkg.go.dev](https://pkg.go.dev/github.com/gin-gonic/gin#New)
   * * [gin/gin.go at v1.7.4 · gin-gonic/gin](https://github.com/gin-gonic/gin/blob/v1.7.4/gin.go#L182)
   * Default()の処理の理解を通じて独自設定を適用するためNew()を採用
  */
  e.Use(CORS())
  /*
   * * [Go+GinでCors設定を行い、クロスオリジンのアクセスを制御する - 親バカエンジニアのナレッジ帳](https://ti-tomo-knowledge.hatenablog.com/entry/2020/06/15/213401)
   * * [gin-contrib/cors： Official CORS gin's middleware](https://github.com/gin-contrib/cors)
   * これらを参考にCORS()を実装
   */
  //ginのログをloggerでとる //フォーマット形式はloggerに依存する //担当：石森
  e.Use(ginzap.Ginzap(logger.Desugar(), time.RFC3339, true))
  //パニック時ステータスコード500を送出 //担当：石森
  e.Use(ginzap.RecoveryWithZap(logger.Desugar(), true))
  /*
   * * [gin-contrib/zap： Alternative logging through zap](https://github.com/gin-contrib/zap#example)
   *   Exampleを参考
   */

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

  
/*
 * * [HTTP status codes](https://golang.org/src/net/http/status.go)
 */

  //#region ユーザエンドポイントの構築
  userRepository := user.NewRepository(db, logger)
  userService := user.NewService(userRepository, logger)
  user.RegisterHandlers(e.Group("/api/user/"), userService, logger)
  //#endregion //担当：石森
  
  //#region プロフィールエンドポイントの構築
  profileRepository := profile.NewRepository(db, logger)
  profileService := profile.NewService(profileRepository, logger)
  profile.RegisterHandlers(e.Group("/api/user/"), profileService, logger)
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
