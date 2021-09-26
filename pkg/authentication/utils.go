package authentication

import (
  "time"
  "github.com/golang-jwt/jwt/v4"
)


const (
  ndaysPerYear = 365
  nhoursPerDay = 24
)


func calcYears2SecondsConversion(nyears int) time.Date {
  return time.Duration(nyears * ndaysPerYear * nhoursPerDay) * time.Hour
}


func NewToken(claims map[string]interface{}, secret string) (string, error) {
  var jwtClaims jwt.MapClaims
  jwtClaims = claims
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
  tokenString, err := token.SignedString([]byte(secret))
  return tokenString, err 
}
