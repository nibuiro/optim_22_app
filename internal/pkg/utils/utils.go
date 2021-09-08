package utils

import (
  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v4"
)


//ユーザIDを取得
func GetUserIdFromHeader(c *gin.Context) string {
  return "dummyID"
}

//アクセストークンからユーザIDを取得
func getUserIdFromHeader(c *gin.Context) string {
  //Authorizationヘッダを取得
  tokenString := c.GetHeader("Authorization")
  //jwtパッケージのParseUnverified()が非公開のため仕方なくParse()を利用するため
  dummySender := func(token *jwt.Token) (interface{}, error) {
    return "", nil
  }
  //トークンをパース 
  token, _ := jwt.Parse(tokenString, dummySender)
  //claimsを辞書型として取得
  claims, _ := token.Claims.(jwt.MapClaims)

  return claims["userID"]
}

