package authentication


import (
//  "net/http"
//  "github.com/gin-gonic/gin"
//  "fmt"
  "github.com/golang-jwt/jwt/v4"
)



type Authorizer struct {
  refreshTokenSecret []byte
  accessTokenSecret []byte
  expiration int
}


func New(refreshTokenSecret string, accessTokenSecret string, expiration int) *Authorizer {
  return &Authorizer{
    refreshTokenSecret: []byte(refreshTokenSecret), 
    accessTokenSecret: []byte(accessTokenSecret), 
    expiration: expiration,
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