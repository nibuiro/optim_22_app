package auth22

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "time"
  "github.com/golang-jwt/jwt/v4"
  "strings"
)


func TestNewToken(t *testing.T) {
  claims := map[string]interface{}{
    "foo": "bar",
    "nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
  }
  token, err := NewToken(claims, []byte("ABC"))
  assert.Nil(t, err)
  assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.Fv2Ff2yy4AUNlZP-p0sG-y5LKwNhIOTpF7ufmDxo0yg", token)
}


func TestValidateSignature(t *testing.T) {
  tests := []struct {
    name      string
    token     string
    wantError bool
  }{
    {
      "valid token", 
      accessToken2021,
      false,
    },
    {
      "invalid token - signed with fake secret", 
      MakeFakeTokenBySignWithFakeSecret(),
      true,
    },
    {
      "invalid token - concat original signing", 
      MakeFakeTokenByConcatOrignalSigning(),
      true,
    },
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      valid, _ := ValidateSignature(tt.token, []byte("secret_key"))
      assert.Equal(t, tt.wantError, !valid)
    })
  }
}


func MakeFakeTokenBySignWithFakeSecret() string {
  //accessToken2100 := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEiLCJleHAiOjQxMjM3MjEwMTd9.JezlU23njkKGldV4ZH1QI37O1yCd0Y-mWmnIu-7aKEo"
  token, _ := jwt.Parse(accessToken2021 , MakeTokenSecretSender([]byte("ABC")))
  claims, _ := token.Claims.(jwt.MapClaims)
  claims["exp"] = CalcFutureUnixTime(365)
  token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, _ := token.SignedString([]byte("ABC"))
  
  return tokenString
}


func MakeFakeTokenByConcatOrignalSigning() string {
  //accessToken2100 := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEiLCJleHAiOjQxMjM3MjEwMTd9.JezlU23njkKGldV4ZH1QI37O1yCd0Y-mWmnIu-7aKEo"
  token, _ := jwt.Parse(accessToken2021 , MakeTokenSecretSender([]byte("ABC")))
  claims, _ := token.Claims.(jwt.MapClaims)
  claims["exp"] = CalcFutureUnixTime(365)
  token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, _ := token.SignedString([]byte("ABC"))
  
  splited_raw := strings.Split(accessToken2021, ".")
  splited_fake := strings.Split(tokenString, ".")
  
  return splited_fake[0]+"."+splited_fake[1]+"."+splited_raw[2]
}
