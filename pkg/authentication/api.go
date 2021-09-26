package authentication


import (
  "net/http"
  "time"
  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v4"
)


const (
  ndaysPerYear = 365
  nhoursPerDay = 24
)


type authInterface struct {
  domain string
  refreshTokenSecret []byte
  accessTokenSecret []byte
  refreshTokenExpiration time.Duration
  accessTokenExpiration time.Duration
  service Service
}


func New(domain string, refreshTokenSecret string, accessTokenSecret string, refreshTokenExpiration int, accessTokenExpiration int, authorizationService AuthorizationService) *authInterface {
  return &authInterface{
    domain: domain,
    refreshTokenSecret: []byte(refreshTokenSecret), 
    accessTokenSecret: []byte(accessTokenSecret), 
    refreshTokenExpiration: time.Duration(refreshTokenExpiration * ndaysPerYear * nhoursPerDay) * time.Hour,
    accessTokenExpiration: time.Duration(accessTokenExpiration * ndaysPerYear * nhoursPerDay) * time.Hour,
  }
}


func (auth *authInterface) RefreshTokenRefreshHandler() gin.HandlerFunc {
  return func(c *gin.Context) {

    if refreshToken, err := c.Cookie("refresh_token"); err != nil {
      c.Status(http.StatusBadRequest)
      return
    } else {
      token, _ := jwt.Parse(refreshToken, auth.refreshTokenSecretSender)
      _, ok := token.Claims.(jwt.MapClaims)
  
      if ok {
        if token.Valid {
          newRefreshToken, _ := auth.service.Refresh(refreshToken)
          c.SetCookie("refresh_token", newRefreshToken, 1, "/", auth.domain, false, true)
            //func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
        } else {
          c.Status(http.StatusUnauthorized)
          return
        }
      } else {
        c.Status(http.StatusBadRequest)
        return
      }
    }
  }
}

//リフレッシュトークンが有効期限内のとき任意のアクセストークンの有効期限を延長
func (auth *authInterface) AccessTokenRefreshHandler() gin.HandlerFunc {
  return func(c *gin.Context) {

    refreshToken, err := c.Cookie("refresh_token")

    if err != nil {
      c.AbortWithStatus(http.StatusBadRequest)
    } else {

      token, _ := jwt.Parse(refreshToken, auth.refreshTokenSecretSender)
      _, ok := token.Claims.(jwt.MapClaims)

      if ok {
        //リフレッシュトークンが期限内
        if token.Valid {

          //#region アクセストークンのclaimsを取り出し有効期限を更新し符号化後、署名

          //Authorizationヘッダーからstring型のトークンを取得
          tokenString := c.GetHeader("Authorization")
          //トークンの改竄と期限を検証
          token, _ := jwt.Parse(tokenString, auth.accessTokenSecretSender)
          //claimsを辞書型として取得
          claims, ok := token.Claims.(jwt.MapClaims)
      
          if ok {
            expiration := time.Now()
            expiration = expiration.Add(auth.accessTokenExpiration)

            claims["exp"] = expiration.Unix()
            newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)              
            // Sign and get the complete encoded token as a string using the secret
            newTokenString, _ := newToken.SignedString([]byte(auth.accessTokenSecret))
            c.Header("Authorization", newTokenString)

            c.Status(http.StatusOK)
          } else {
            c.Status(http.StatusBadRequest)
            return
          }
          //#endregion
        } else {
          c.Status(http.StatusUnauthorized)
          return
        }
      } else {
        c.Status(http.StatusBadRequest)
        return
      }
    }
  }
}

//認証情報を空文字列で上書き
func (auth *authInterface) RevokeHandler() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.SetCookie("refresh_token", "", 0, "/", "", false, true)
    c.Header("Authorization", "")
  }
}

//パース関数にリフレッシュトークン用秘密鍵を渡すコールバック
func (auth *authInterface) refreshTokenSecretSender(token *jwt.Token) (interface{}, error) {
  return auth.refreshTokenSecret, nil
}

//パース関数にアクセストークン用秘密鍵を渡すコールバック
func (auth *authInterface) accessTokenSecretSender(token *jwt.Token) (interface{}, error) {
  return auth.accessTokenSecret, nil
}
