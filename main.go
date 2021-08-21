package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  // コメントを外す
  // "optim_22_app/model"
)

func main() {
  // DB接続後、マイグレーションを実行する。
  // 手順としては、まずコンテナを立ち上げた後、mysqlでoptim_devデータベースを作成する。
  // その後、model.InitDB(),import(optim_22_app/model)のコメントを外し、カレントディレクトリでgo run main.goを実行する。
  // プログラムの詳細はmodel/migrate.goに記載。
  // model.InitDB()

  // ルーターを作成している
  router := gin.Default()
  // helloメソッドがwebブラウザから指定された場合、Hello World!!を返す。
  // gin.Contextはリクエストとレスポンスのやり取りをするための型である。
  // リクエストがokの場合、Hello World!!を返す。
  router.GET("/hello", func(c *gin.Context) {
    c.String(http.StatusOK, "Hello World!!")
  })
  // 8080ポートで実行。
  router.Run(":8080")
}
