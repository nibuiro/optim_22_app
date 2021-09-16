package profile

import (
  "github.com/gin-gonic/gin"
  //"net/http"
  "optim_22_app/pkg/log"
  "optim_22_app/internal/pkg/config"
  //"optim_22_app/internal/pkg/utils"

)

//プロフィール操作の依存関係
type resource struct {
  config *config.Config
  service Service
  logger  log.Logger
}

//プロフィール操作についてエンドポイントを登録
func RegisterHandlers(r *gin.RouterGroup, config *config.Config, service Service, logger log.Logger) {

  rc := resource{config, service, logger}

  //取得
  r.GET("/api/profile", rc.get())
  //登録
  r.POST("/api/profile", rc.post())
  //修正
  r.PATCH("/api/profile", rc.patch())
  //削除
  r.DELETE("/api/profile", rc.delete())

}


func (rc resource) get() gin.HandlerFunc {
  return func(c *gin.Context) {
  }
}


func (rc resource) post() gin.HandlerFunc {
  return func(c *gin.Context) {
  }
}


func (rc resource) patch() gin.HandlerFunc {
  return func(c *gin.Context) {
  }
}


func (rc resource) delete() gin.HandlerFunc {
  return func(c *gin.Context) {
  }
}





