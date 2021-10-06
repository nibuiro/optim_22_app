package auth22


import (
  //"io"
  "net/http"
 // "time"
  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v4"
)

//（メソッド名、パス名）の順列を保持
type Rule map[string]map[string]bool


type resource struct {
  service Service
  domain string
}

//認証APIを発行
func New(service Service, domain string) *resource {
  return &resource{
    service: service,
    domain: domain,
  }
}

//リフレッシュトークンが有効期限内のときアクセストークンの有効期限を延長
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

//リフレッシュトークンが有効期限内のときアクセストークンの有効期限を延長
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
            if !isValid {
              //リフレッシュトークンを利用したアクセストークンの不正取得の試行
              c.Status(http.StatusBadRequest)
              return
            } else {
              if newTokenString, err := rc.service.RefreshAccessToken(claims); err != nil {
                c.Status(http.StatusBadRequest)
                return
              } else {
                SetTokenWithControl(c, "", newTokenString)
                c.Status(http.StatusOK)
                return
              }
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

    if body, err := c.GetRawData(); err != nil {
      c.Status(http.StatusBadRequest)
      return
    } else {
      if credential, err := rc.service.ReadCredential(body); err != nil {
        c.Status(http.StatusBadRequest)
        return
      } else {
        //資格情報の確認
        if valid, err := rc.service.ValidateCredential(c.Request.Context(), claims, credential); err != nil {
          c.Status(http.StatusBadRequest)
          return 
        } else {
          if !valid {
            //不正なログイン
            c.Status(http.StatusUnauthorized)
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
}

//リフレッシュトークンが有効期限内のときリフレッシュトークンとアクセストークンの有効期限を延長
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
        } else {
          //認証情報取得
          if refreshToken, err := rc.service.RefreshRefreshToken(claims); err != nil {
            c.Status(http.StatusInternalServerError)
            return
          } else {
            //リフレッシュトークンが期限内
            claims = make(jwt.MapClaims)
            accessToken := c.GetHeader("Authorization")
            if _, err := rc.service.ReadAccessToken(claims, accessToken); err != nil {
              c.Status(http.StatusBadRequest)
              return
            } else {
              if isValid, err := rc.service.ValidateAccessTokenSignature(accessToken); err != nil {
                c.Status(http.StatusInternalServerError)
                return
              } else {
                if !isValid {
                  //リフレッシュトークンを利用したアクセストークンの不正取得の試行
                  c.Status(http.StatusBadRequest)
                  return
                } else {
                  if accessToken, err := rc.service.RefreshAccessToken(claims); err != nil {
                    c.Status(http.StatusBadRequest)
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

//カスタムヘッダー登録許可申請
func SetTokenWithControl(c *gin.Context, refreshToken string, accessToken string) {
  c.Header("Authorization", accessToken)
  c.Header("Refresh-Token", refreshToken)
  c.Header("Access-Control-Allow-Origin", "*")
  c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,HEAD,OPTION")
  c.Header("Access-Control-Request-Headers", "Authorization,Refresh-Token")
  c.Header("Access-Control-Expose-Headers", "Authorization,Refresh-Token")
}