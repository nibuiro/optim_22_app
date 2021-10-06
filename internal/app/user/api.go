package user

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "optim_22_app/pkg/log"
  "optim_22_app/internal/pkg/utils"
  "encoding/json"
  "optim_22_app/internal/pkg/auth22"
  "bytes"

)

//ユーザ操作の依存関係
type resource struct {
  service Service
  logger  log.Logger
}

//ユーザ操作についてエンドポイントを登録
func RegisterHandlers(r *gin.RouterGroup, service Service, logger log.Logger) {

  rc := resource{service, logger}

  //登録する
  r.POST("/api/user", rc.post())
  //退会する
  r.DELETE("/api/user", rc.delete())

}

func (rc resource) post() gin.HandlerFunc {
  return func(c *gin.Context) {
    var input RegistrationInformation
  
    //BodyからJSONをパースして読み取る
    if err := c.BindJSON(&input); err != nil {
      rc.logger.Error(err)
      c.Status(http.StatusBadRequest)
      return 
    }
  
    //ユーザ作成及び認証情報取得 
    _, err := rc.service.Create(c.Request.Context(), input)
    if err != nil {
      rc.logger.Error(err)
      c.Status(http.StatusBadRequest)
      return 
    } else {
      reader := auth22.Credential{
        Email: input.Email,
        Password: input.Password,
      }
      credentialJSON, err := json.Marshal(reader)
      if err != nil {
        rc.logger.Debug(err, string(credentialJSON))
        c.Status(http.StatusInternalServerError)
        return
      } else {
        rc.logger.Debug(string(credentialJSON))
        c.Writer = &utils.HttpBodyWriter{
          Body: bytes.NewBuffer(credentialJSON), 
          ResponseWriter: c.Writer,
        }
        c.Redirect(http.StatusTemporaryRedirect, "/auth")
        return
      }
    }
  }
}


func (rc resource) delete() gin.HandlerFunc {
  return func(c *gin.Context) {
    userId := utils.GetUserIdFromHeaderAsInt(c)
    err := rc.service.Delete(c.Request.Context(), userId)
    if err != nil {
      rc.logger.Error(err)
      c.Status(http.StatusBadRequest)
      return 
    } else {
      c.Status(http.StatusOK)
      return 
    }
  }
}





