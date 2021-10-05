package auth22

import (
  "time"
  "github.com/golang-jwt/jwt/v4"
  "github.com/gin-gonic/gin"
 // "errors"
  "strconv"
)


const (
  ndaysPerYear = 365
  nhoursPerDay = 24
)



func CalcFutureUnixTime(ndays int) int64 {
    expiration := time.Now()
    nhours := time.Duration(ndays * nhoursPerDay) //単位は10^-9秒
    expiration = expiration.Add(nhours * time.Hour)
    return expiration.Unix()
}


func NewToken(claims map[string]interface{}, secret []byte) (string, error) {
  var jwtClaims jwt.MapClaims
  jwtClaims = claims
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
  tokenString, err := token.SignedString(secret)
  return tokenString, err 
}

func ValidateSignature(tokenString string, secret []byte) (bool, error) {
  token, _ := jwt.Parse(tokenString, MakeTokenSecretSender(secret))

  if verificationTokenString, err := token.SignedString(secret); err != nil {
    return false, err
  } else {
    isSame := (tokenString == verificationTokenString)
    return isSame, nil
  }

}


func MakeTokenSecretSender(secret[]byte) func(token *jwt.Token) (interface{}, error) {
  secretSender := func(token *jwt.Token) (interface{}, error) {
    return secret, nil
  }
  return secretSender
}




func SetTokenWithControl(c *gin.Context, refreshToken string, accessToken string) {
  c.Header("Authorization", accessToken)
  c.Header("Refresh-Token", refreshToken)
  c.Header("Access-Control-Allow-Origin", "*")
  c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,HEAD,OPTION")
  c.Header("Access-Control-Request-Headers", "Authorization,Refresh-Token")
  c.Header("Access-Control-Expose-Headers", "Authorization,Refresh-Token")
}
