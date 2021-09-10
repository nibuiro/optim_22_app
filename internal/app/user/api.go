package user

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "optim_22_app/pkg/log"
  "optim_22_app/internal/pkg/config"
  "optim_22_app/internal/pkg/utils"

)

//ユーザ操作の依存関係
type resource struct {
  config *config.Config
  service Service
  logger  log.Logger
}

//ユーザ操作についてエンドポイントを登録
func RegisterHandlers(r *gin.RouterGroup, config *config.Config, service Service, logger log.Logger) {

  rc := resource{config, service, logger}

  //登録する
  r.POST("/api/user", rc.create())
  //退会する
  r.DELETE("/api/user", rc.delete())
  //ログイン
  r.POST("/api/session", rc.login())
  //ログアウト
  r.DELETE("/api/session", rc.logout())

}


func (rc resource) create() gin.HandlerFunc {
  return func(c *gin.Context) {
    var input registrationInformation
  
    //BodyからJSONをパースして読み取る
    if err := c.BindJSON(&input); err != nil {
      rc.logger.Error(err)
    }
  
    //ユーザ作成及び認証情報取得
    refreshToken, accessToken, err := rc.service.Create(c.Request.Context(), input)
    if err != nil {
      rc.logger.Error(err)
    }
    
    //#region ヘッダに認証情報を付加
    c.Header("Authorization", accessToken)
    c.SetCookie("refresh_token", refreshToken, 1, "/", rc.config.Domain, false, true)
    c.Status(http.StatusCreated)
    //#endregion
  }
}


func (rc resource) delete() gin.HandlerFunc {
  return func(c *gin.Context) {
    userId := utils.GetUserIdFromHeaderAsInt(c)
    err := rc.service.Delete(c.Request.Context(), userId)
    if err != nil {
      rc.logger.Error(err)
    }
  }
}


func (rc resource) login() gin.HandlerFunc {
  return func(c *gin.Context) {

    var input loginInformation
  
    //BodyからJSONをパースして読み取る
    if err := c.BindJSON(&input); err != nil {
      rc.logger.Error(err)
    }
  
    //資格情報確認及び認証情報取得
    refreshToken, accessToken, err := rc.service.Login(c.Request.Context(), input)
    if err != nil {
      rc.logger.Error(err)
    }
    
    //#region ヘッダに認証情報を付加
    c.Header("Authorization", accessToken)
    c.SetCookie("refresh_token", refreshToken, 1, "/",  rc.config.Domain, false, true)
    c.Status(http.StatusCreated)
    //#endregion
  }
}


func (rc resource) logout() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.SetCookie("refresh_token", "", 0, "/", "", false, true)
    c.Status(http.StatusOK)
  }
}



