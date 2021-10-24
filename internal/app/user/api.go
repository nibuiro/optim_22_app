package user

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "optim_22_app/pkg/log"
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

}

func (rc resource) post() gin.HandlerFunc {
  return func(c *gin.Context) {
    var input RegistrationInformation
  
    //BodyからJSONをパースして読み取る
    if err := c.ShouldBindJSON(&input); err != nil {
      rc.logger.Error(err)
      c.Status(http.StatusBadRequest)
      return 
    }
    rc.logger.Debug(input)
  
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
        //c.Redirect(http.StatusTemporaryRedirect, "/auth")
        
        //トークン取得
        resp, err := http.Post("http://localhost:8080/auth", "application/json", bytes.NewBuffer(credentialJSON))
        defer resp.Body.Close()
        if err != nil {
          rc.logger.Debug(err)
          c.Status(http.StatusInternalServerError)
          return
        } else {
          //rc.logger.Debug(resp.Header.Get("Authorization"))
          //rc.logger.Debug(resp.Header.Get("Refresh-Token"))
          c.Header("Authorization", resp.Header.Get("Authorization"))
          c.Header("Refresh-Token", resp.Header.Get("Refresh-Token"))
          //CORSによる設定のため不要？
          c.Header("Access-Control-Allow-Origin", "*")
          c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,HEAD,OPTION")
          c.Header("Access-Control-Request-Headers", "Authorization,Refresh-Token")
          c.Header("Access-Control-Expose-Headers", "Authorization,Refresh-Token")
          //
          c.Status(http.StatusOK)
          //body, err := io.ReadAll(resp.Body) //不要
          return
        }
      }
    }
  }
}
