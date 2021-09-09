package authentication


import (
//  "net/http"
//  "github.com/gin-gonic/gin"
//  "fmt"
//  "github.com/golang-jwt/jwt/v4"
  "time"
)


const (
  ndaysPerYear = 365
  nhoursPerDay = 24
)


type AuthorizationService interface {
  New(args ...interface{}) string
  Refresh(refreshToken string) string
}


type Authorizer struct {
  domain string
  refreshTokenSecret []byte
  accessTokenSecret []byte
  validityPeriod time.Duration
  authorizationService AuthorizationService
}


func New(domain string, refreshTokenSecret string, accessTokenSecret string, validityPeriod int, authorizationService AuthorizationService) *Authorizer {
  return &Authorizer{
    domain: domain,
    refreshTokenSecret: []byte(refreshTokenSecret), 
    accessTokenSecret: []byte(accessTokenSecret), 
    validityPeriod: time.Duration(validityPeriod * ndaysPerYear * nhoursPerDay) * time.Hour,
    authorizationService: authorizationService,
  }
}

