package auth22


import (
  //"io"
  "net/http"
 // "time"
  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v4"
  //"fmt"
  "optim_22_app/pkg/log"
)


type Rule map[string]map[string]bool


type resource struct {
  service Service
  logger log.Logger
  domain string
}


func New(service Service, logger log.Logger, domain string) *resource {
  return &resource{
    service: service,
    logger: logger,
    domain: domain,
  }
}


func (rc *resource) RefreshTokenRefreshHandler() gin.HandlerFunc {
  return func(c *gin.Context) {

    claims := make(jwt.MapClaims)

    if refreshToken := c.GetHeader("Refresh-Token"); refreshToken == "" {
      c.Status(http.StatusBadRequest)
      return
    } else {
      if valid, err := rc.service.ReadRefreshToken(claims, refreshToken); err != nil {
        c.Status(http.StatusBadRequest)
        return
      } else {
        if !valid {
          c.Status(http.StatusUnauthorized)
          return
        }
        if newRefreshToken, err := rc.service.RefreshRefreshToken(claims); err != nil {
          c.Status(http.StatusInternalServerError)
          return
        } else {
          SetTokenWithControl(c, newRefreshToken, "")
          c.Status(http.StatusOK)
          return
        }
      }
    }
  }
}

//リフレッシュトークンが有効期限内のとき任意のアクセストークンの有効期限を延長
func (rc *resource) AccessTokenRefreshHandler() gin.HandlerFunc {
  return func(c *gin.Context) {

    claims := make(jwt.MapClaims)

    refreshToken := c.GetHeader("Refresh-Token")

    if refreshToken == "" {
      c.Status(http.StatusBadRequest)
      return
    } else {

      if valid, err := rc.service.ReadRefreshToken(claims, refreshToken); err != nil {
        c.Status(http.StatusBadRequest)
        return
      } else {
        if !valid {
          c.Status(http.StatusUnauthorized)
          return
        }
        //リフレッシュトークンが期限内
        accessToken := c.GetHeader("Authorization")
        if _, err := rc.service.ReadAccessToken(claims, accessToken); err != nil {
          c.Status(http.StatusBadRequest)
          return
        } else {
          if isValid, err := rc.service.ValidateAccessTokenSignature(accessToken); err != nil {
            c.Status(http.StatusInternalServerError)
            return
          } else {
            if isValid {
              if newTokenString, err := rc.service.RefreshAccessToken(claims); err != nil {
                c.Status(http.StatusBadRequest)
                return
              } else {
                SetTokenWithControl(c, "", newTokenString)
                c.Status(http.StatusOK)
                return
              }
            } else {
              //リフレッシュトークンを利用したアクセストークンの不正取得の試行
              c.Status(http.StatusBadRequest)
              return
            }
          }
        }
      }
    }
  }
}

//Validate credential and get both access token and refresh token.
func (rc resource) Login() gin.HandlerFunc {
  return func(c *gin.Context) {

    claims := make(jwt.MapClaims)
  
    //buf := make([]byte, 1028)
    
    //BodyからJSONをパースして読み取る
      //  rc.logger.Debug(buf)
    if body, err := c.GetRawData(); err != nil {
      c.Status(http.StatusBadRequest)
      return
    } else {
      rc.logger.Debug(string(body))
      if credential, err := rc.service.ReadCredential(body); err != nil {
        rc.logger.Debug(err)
        c.Status(http.StatusBadRequest)
        return
      } else {
        rc.logger.Debug(credential)
        //資格情報の確認
        if err := rc.service.ValidateCredential(c.Request.Context(), claims, credential); err != nil {
          c.Status(http.StatusBadRequest)
          rc.logger.Debug(err)
          return 
        } else {

          //認証情報取得
          if refreshToken, err := rc.service.GenerateRefreshToken(claims); err != nil {
            c.Status(http.StatusInternalServerError)
            return
          } else {
            if accessToken, err := rc.service.GenerateAccessToken(claims); err != nil {
              c.Status(http.StatusInternalServerError)
              return
            } else {
              SetTokenWithControl(c, refreshToken, accessToken)
              c.Status(http.StatusOK)
              return 
            }
          }
        }
      }
    }
  }
}


func (rc resource) RefreshAccessTokenAndRefreshToken() gin.HandlerFunc {
  return func(c *gin.Context) {

    claims := make(jwt.MapClaims)

    refreshToken := c.GetHeader("Refresh-Token")

    if refreshToken == "" {
      c.Status(http.StatusBadRequest)
      return
    } else {  
      if valid, err := rc.service.ReadRefreshToken(claims, refreshToken); err != nil {
        c.Status(http.StatusBadRequest)
        return
      } else {
        if !valid {
          c.Status(http.StatusBadRequest)
          return
        }
        //認証情報取得
        if refreshToken, err := rc.service.GenerateRefreshToken(claims); err != nil {
          c.Status(http.StatusInternalServerError)
          return
        } else {
          if accessToken, err := rc.service.RefreshAccessToken(claims); err != nil {
            c.Status(http.StatusInternalServerError)
            return
          } else {
            //#region ヘッダに認証情報を付加
            SetTokenWithControl(c, refreshToken, accessToken)
            c.Status(http.StatusOK)
            //#endregion
            return 
          }
        }
      }
    }
  }
}

//認証情報を空文字列で上書き
func (rc *resource) RevokeHandler() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.SetCookie("refresh_token", "", 0, "/", "", false, true)
    c.Header("Authorization", "")
  }
}
