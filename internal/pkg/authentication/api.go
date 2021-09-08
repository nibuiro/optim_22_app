package authentication


import (
//  "net/http"
  "github.com/gin-gonic/gin"
//  "fmt"
  "github.com/golang-jwt/jwt/v4"
)

const (
  ndaysPerYear = 365
  nhoursPerDay = 24
)


type Authorizer struct {
  refreshTokenSecret []byte
  accessTokenSecret []byte
  validityPeriod int
  authorizationService interface{ Endpoint func(*gin.Context), }
}


func New(refreshTokenSecret string, accessTokenSecret string, validityPeriod int, authorizationService interface{}) *Authorizer {
  return &Authorizer{
    refreshTokenSecret: []byte(refreshTokenSecret), 
    accessTokenSecret: []byte(accessTokenSecret), 
    validityPeriod: validityPeriod * ndaysPerYear * nhoursPerDay,
    authorizationService: authorizationService,
  }
}



//パース関数にリフレッシュトークン用秘密鍵を渡すコールバック
func (auth *Authorizer) refreshTokenSecretSender(token *jwt.Token) (interface{}, error) {
  return auth.refreshTokenSecret, nil
}

//パース関数にアクセストークン用秘密鍵を渡すコールバック
func (auth *Authorizer) accessTokenSecretSender(token *jwt.Token) (interface{}, error) {
  return auth.accessTokenSecret, nil
}