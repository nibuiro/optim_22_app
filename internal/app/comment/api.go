package comment

import (
  "net/http"
  "strconv"
  "encoding/json"
  "github.com/gin-gonic/gin"
  "optim_22_app/pkg/log"
  "optim_22_app/internal/pkg/config"
 // b64 "encoding/base64"
//  "optim_22_app/internal/pkg/utils"
//  "reflect"
)

//コメント操作の依存関係
type resource struct {
  config *config.Config
  service Service
  logger  log.Logger
}

//コメント操作についてエンドポイントを登録
func RegisterHandlers(r *gin.RouterGroup, config *config.Config, service Service, logger log.Logger) {

  rc := resource{config, service, logger}

  //ディスカッション ID(:requestID) のコメント一覧を取得
  r.GET("/api/discussion/:requestID", rc.get())
  //ディスカッション ID(:requestID) にコメントを投稿
  r.POST("/api/discussion/:requestID", rc.createStub())
  //ディスカッション ID(:requestID) に投稿されている
  //コメント ID(:commentID) を削除
  r.DELETE("/api/discussion/:requestID/:commentID", rc.deleteStub())

}


func (rc resource) get() gin.HandlerFunc {
  return func(c *gin.Context) {
    requestID := c.Param("requestID")

    if !isIntegerString(requestID) {
      c.Status(http.StatusBadRequest)
      return
    }

    comments, err := rc.service.Get(c.Request.Context(), requestID)
    //rc.logger.Debug(comments)
    if err != nil {
      rc.logger.Error(err)
      c.Status(http.StatusNotFound)
      return 
    } else {
      if commentsText, err := json.Marshal(comments); err != nil {
        rc.logger.Error(err)
        c.Status(http.StatusInternalServerError)
        return
      } else {
        //rc.logger.Debug(reflect.TypeOf(commentsText))
        c.Header("Content-Type", "application/json")
        c.String(http.StatusOK, string(commentsText[:]))
        return
      }
    }
  }
}


func isIntegerString(query string) bool {
  _, err := strconv.Atoi(query)
  return err == nil
}