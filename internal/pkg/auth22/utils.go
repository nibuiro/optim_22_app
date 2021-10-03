package auth22

import (
  "time"
  "github.com/golang-jwt/jwt/v4"
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