package user

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "optim_22_app/pkg/log"
)

//`POST /api/user`が要求する情報
type registrationInformation struct {
  name     string `json:"name"`
  email    string `json:"email"`
  password string `json:"password"`
}

//ユーザ操作の依存関係
type resource struct {
  service Service
  logger  log.Logger
}

//ユーザ操作についてエンドポイントを登録
func RegisterHandlers(r *gin.RouterGroup, service Service, logger log.Logger) {

  res := resource{service, logger}

  //登録する
  r.POST("/api/user", res.create)
  //退会する
  r.DELETE("/api/user", res.delete)
  //ログイン
  r.POST("/api/session", res.login)
  //ログアウト
  r.DELETE("/api/session", res.logout)

}