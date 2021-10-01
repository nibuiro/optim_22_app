package authentication

import (
  "net/http"
//  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v4"
)

func (auth *Authorizer) ValidateAccessToken() gin.HandlerFunc {
  return func(c *gin.Context) {

    //Authorizationヘッダーからstring型のトークンを取得
  	tokenString := c.GetHeader("Authorization")
  	//トークンの改竄と期限を検証
  	//tips: expキーがない場合無期限トークンとして扱われ、token.Validの値はtrue
    token, _ := jwt.Parse(tokenString, auth.accessTokenSecretSender)
    //辞書型に変換
    _, ok := token.Claims.(jwt.MapClaims)
    //claims, ok := token.Claims.(jwt.MapClaims)

    if ok {
      if token.Valid {
        //なにもしない
      } else {
        c.AbortWithStatus(http.StatusUnauthorized)
      }
    } else {
      c.AbortWithStatus(http.StatusBadRequest)
    }

    c.Next()
  }
}