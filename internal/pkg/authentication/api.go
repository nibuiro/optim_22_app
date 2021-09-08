package authentication


import (
//  "net/http"
  "github.com/gin-gonic/gin"
//  "fmt"
//  "github.com/golang-jwt/jwt/v4"
  "time"
)


const (
  ndaysPerYear = 365
  nhoursPerDay = 24
)


type AuthorizationService interface {
  Endpoint(c *gin.Context)
}


type Authorizer struct {
  refreshTokenSecret []byte
  accessTokenSecret []byte
  validityPeriod time.Duration
  authorizationService AuthorizationService
}


func New(refreshTokenSecret string, accessTokenSecret string, validityPeriod int, authorizationService AuthorizationService) *Authorizer {
  return &Authorizer{
    refreshTokenSecret: []byte(refreshTokenSecret), 
    accessTokenSecret: []byte(accessTokenSecret), 
    validityPeriod: time.Duration(validityPeriod * ndaysPerYear * nhoursPerDay) * time.Hour,
    authorizationService: authorizationService,
  }
}

