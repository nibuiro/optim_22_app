package authentication

import (
  "net/http"
  "time"
  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v4"
)

func (auth *Authorizer) refreshTokenRefreshHandler() gin.HandlerFunc {
  return func(c *gin.Context) {

    refreshToken, err := c.Cookie("refresh_token")

    if err != nil {
      c.AbortWithStatus(http.StatusBadRequest)
    } else {

      token, _ := jwt.Parse(refreshToken, auth.refreshTokenSecretSender)
      _, ok := token.Claims.(jwt.MapClaims)
  
      if ok {
        if token.Valid {
          //なにもしない
        } else {
          c.AbortWithStatus(http.StatusUnauthorized)
        }
      } else {
        c.AbortWithStatus(http.StatusBadRequest)
      }
    }
  }
}

//リフレッシュトークンが有効期限内のとき任意のアクセストークンの有効期限を延長
func (auth *Authorizer) accessTokenRefreshHandler() gin.HandlerFunc {
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
            expiration = expiration.Add(time.Duration(auth.validityPeriod) * time.Hour)

            claims["exp"] = expiration.Unix()
            new_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)              
            // Sign and get the complete encoded token as a string using the secret
            new_tokenString, _ := new_token.SignedString([]byte(auth.accessTokenSecret))
            c.Header("Authorization", new_tokenString)

            c.Status(http.StatusOK)
          } else {
            c.AbortWithStatus(http.StatusBadRequest)
          }

          //#endregion
        } else {
          
          c.AbortWithStatus(http.StatusUnauthorized)
        }
      } else {
        c.AbortWithStatus(http.StatusBadRequest)
      }
    }
  }
}



func (auth *Authorizer) revokeHandler() gin.HandlerFunc {
  return func(c *gin.Context) {

    c.AbortWithStatus(http.StatusBadRequest)
  }
}
