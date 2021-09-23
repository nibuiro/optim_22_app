package comment

import (
//  "net/http"
  "github.com/gin-gonic/gin"
  "optim_22_app/pkg/log"
  "optim_22_app/internal/pkg/config"
//  "optim_22_app/internal/pkg/utils"

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
  r.GET("/api/discussion/:requestID", rc.getStub())
  //ディスカッション ID(:requestID) にコメントを投稿
  r.POST("/api/discussion/:requestID", rc.createStub())
  //ディスカッション ID(:requestID) に投稿されている
  //コメント ID(:commentID) を削除
  r.DELETE("/api/discussion/:requestID/:commentID", rc.deleteStub())

}
