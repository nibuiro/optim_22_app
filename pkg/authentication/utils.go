package authentication

import (
  "github.com/golang-jwt/jwt/v4"
//  "fmt"
)

func NewToken(claims map[string]interface{}, secret string) (string, error) {
  var jwtClaims jwt.MapClaims
  jwtClaims = claims
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
  tokenString, err := token.SignedString([]byte(secret))
  return tokenString, err 
}