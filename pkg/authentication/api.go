package authentication


import (
  "io"
  "net/http"
 // "time"
  "github.com/gin-gonic/gin"
)


type Rule map[string]map[string]bool


type resource struct {
  service Service
  domain string
}


func New(service Service, domain string) *resource {
  return &resource{
    service: service,
    domain: domain,
  }
}


func (rc *resource) RefreshTokenRefreshHandler() gin.HandlerFunc {
  return func(c *gin.Context) {

    s := rc.service
    s.RefreshService(c.Request.Context())

    if refreshToken, err := c.Cookie("refresh_token"); err != nil {
      c.Status(http.StatusBadRequest)
      return
    } else {
      if valid, err := s.ReadRefreshToken(refreshToken); err != nil {
        c.Status(http.StatusBadRequest)
        return
      } else {
        if !valid {
          c.Status(http.StatusUnauthorized)
          return
        }
        if newRefreshToken, err := s.RefreshRefreshToken(); err != nil {
          c.Status(http.StatusUnauthorized)
          return
        } else {
          c.SetCookie("refresh_token", newRefreshToken, 1, "/", rc.domain, false, true)  
          return
        }
      }
    }
  }
}

//リフレッシュトークンが有効期限内のとき任意のアクセストークンの有効期限を延長
func (rc *resource) AccessTokenRefreshHandler() gin.HandlerFunc {
  return func(c *gin.Context) {

    s := rc.service
    s.RefreshService(c.Request.Context())

    refreshToken, err := c.Cookie("refresh_token")

    if err != nil {
      c.AbortWithStatus(http.StatusBadRequest)
    } else {

      if valid, err := s.ReadRefreshToken(refreshToken); err != nil {
        c.Status(http.StatusBadRequest)
        return
      } else {
        if !valid {
          c.Status(http.StatusUnauthorized)
          return
        }
        //リフレッシュトークンが期限内
        accessToken := c.GetHeader("Authorization")
      
        if _, err := s.ReadAccessToken(accessToken); err != nil {
          c.Status(http.StatusBadRequest)
          return
        } else {
/*
          expiration := time.Now()
          expiration = expiration.Add(rc.accessTokenExpiration)

          claims["exp"] = expiration.Unix()
          newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)              
          // Sign and get the complete encoded token as a string using the secret
          newTokenString, _ := newToken.SignedString([]byte(rc.accessTokenSecret))
*/
          if newTokenString, err := s.RefreshAccessToken(); err != nil {
            c.Status(http.StatusUnauthorized)
            return
          } else {
            c.Header("Authorization", newTokenString)
            c.Status(http.StatusOK)
            return
          }
        }
      }
    }
  }
}

//Validate credential and get both access token and refresh token.
func (rc resource) Login() gin.HandlerFunc {
  return func(c *gin.Context) {

    s := rc.service
    s.RefreshService(c.Request.Context())
  
    //BodyからJSONをパースして読み取る
    if body, err := io.ReadAll(c.Request.Body); err != nil {
      c.Status(http.StatusBadRequest)
      return
    } else {
      if err := s.ReadCredential(body); err != nil {
        c.Status(http.StatusBadRequest)
        return
      } else {
        //pass
      }
    }
    //資格情報の確認
    if err := s.ValidateCredential(); err != nil {
      c.Status(http.StatusUnauthorized)
      return 
    } else {
      //認証情報取得
      if refreshToken, err := s.GenerateRefreshToken(); err != nil {
        c.Status(http.StatusInternalServerError)
        return
      } else {
        if accessToken, err := s.RefreshAccessToken(); err != nil {
          c.Status(http.StatusInternalServerError)
          return
        } else {
          //#region ヘッダに認証情報を付加
          c.Header("Authorization", accessToken)
          c.SetCookie("refresh_token", refreshToken, 1, "/",  rc.domain, false, true)
          c.Status(http.StatusCreated)
          //#endregion
          return 
        }
      }
    }
  }
}


func (rc resource) RefreshAccessTokenAndRefreshToken() gin.HandlerFunc {
  return func(c *gin.Context) {

    s := rc.service
    s.RefreshService(c.Request.Context())

    refreshToken, err := c.Cookie("refresh_token")

    if err != nil {
      c.Status(http.StatusBadRequest)
      return
    } else {  
      if valid, err := s.ReadRefreshToken(refreshToken); err != nil {
        c.Status(http.StatusBadRequest)
        return
      } else {
        if !valid {
          c.Status(http.StatusUnauthorized)
          return
        }
        //認証情報取得
        if refreshToken, err := s.GenerateRefreshToken(); err != nil {
          c.Status(http.StatusInternalServerError)
          return
        } else {
          if accessToken, err := s.RefreshAccessToken(); err != nil {
            c.Status(http.StatusInternalServerError)
            return
          } else {
            //#region ヘッダに認証情報を付加
            c.Header("Authorization", accessToken)
            c.SetCookie("refresh_token", refreshToken, 1, "/",  rc.domain, false, true)
            c.Status(http.StatusCreated)
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
