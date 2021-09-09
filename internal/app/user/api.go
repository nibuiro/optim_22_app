package user

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "optim_22_app/pkg/log"
  "optim_22_app/internal/pkg/Config"
)

//`POST /api/user`が要求する情報
type registrationInformation struct {
  name     string `json:"name"`
  email    string `json:"email"`
  password string `json:"password"`
}

//ユーザ操作の依存関係
type resource struct {
  config Config
  service Service
  logger  log.Logger
}

//ユーザ操作についてエンドポイントを登録
func RegisterHandlers(r *gin.RouterGroup, service Service, logger log.Logger) {

  rc := resource{service, logger}

  //登録する
  r.POST("/api/user", rc.create)
  //退会する
  r.DELETE("/api/user", rc.delete)
  //ログイン
  r.POST("/api/session", rc.login)
  //ログアウト
  r.DELETE("/api/session", rc.logout)

}


func (rc resource) create(c *gin.Context) error {

  var input registrationInformation

  //BOdyからJSONをパースして読み取る
  if err := c.BindJSON(&input); err != nil {
    return err
  }

  //ユーザ作成及び認証情報取得
  refreshToken, accessToken, err := rc.service.Create(c.Request.Context(), input)
  if err != nil {
    return err
  }
  
  //#region ヘッダに認証情報を付加
  c.Header("Authorization", accessToken)
  c.SetCookie(name, refreshToken, 1, "/",  rc.config.domain, false, true)
  c.Status(http.StatusCreated)
  //#endregion
}



