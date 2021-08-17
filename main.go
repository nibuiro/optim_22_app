package main
import (
  "net/http"
  "github.com/gin-gonic/gin"
)
func main() {
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
