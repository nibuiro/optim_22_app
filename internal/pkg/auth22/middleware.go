package auth22

import (
  "net/http"
//  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v4"
)

type Ruler func(string, string) bool

func (rc *resource) ValidateAccessToken(rule Rule, methodFirst bool) gin.HandlerFunc {

  IsAllowed := makeRuler(rule, methodFirst)

  return func(c *gin.Context) {

    claims := make(jwt.MapClaims)

    if !IsAllowed(c.Request.Method, c.FullPath()) {
      //Authorizationヘッダーからstring型のトークンを取得
      tokenString := c.GetHeader("Authorization")

      if hasToken := "" != tokenString; !hasToken {
        c.AbortWithStatus(http.StatusForbidden)
      } else {
        if valid, err := rc.service.ReadAccessToken(claims, tokenString); err != nil {
          c.AbortWithStatus(http.StatusBadRequest)
        } else {
          if valid {
            //なにもしない
          } else {
            c.AbortWithStatus(http.StatusForbidden)
          }
        }
      }
    }

    c.Next()
  }
}


func makeRuler(rule Rule, methodFirst bool) Ruler {

  isAllowedMethodAndAllowedEndpoint := func (method string, endpoint string) bool {
    if allowedEndpointSet := rule["*"]; (allowedEndpointSet != nil) {
      if isAllowedEndpoint := allowedEndpointSet["*"]; isAllowedEndpoint {
        return true
      } else if isAllowedEndpoint := allowedEndpointSet[endpoint]; isAllowedEndpoint {
        return true
      } else {
        return false
      }
    } else if allowedEndpointSet := rule[method]; (allowedEndpointSet != nil) {
      if isAllowedEndpoint := allowedEndpointSet["*"]; isAllowedEndpoint {
        return true
      } else if isAllowedEndpoint := allowedEndpointSet[endpoint]; isAllowedEndpoint {
        return true
      } else {
        return false
      }
    } else {
      return false
    }
  }
  
  isAllowedEndpointAndAllowedMethod := func (method string, endpoint string) bool {
    if allowedMethodSet := rule["*"]; (allowedMethodSet != nil) {
      if isAllowedMethod := allowedMethodSet["*"]; isAllowedMethod {
        return true
      } else if isAllowedMethod := allowedMethodSet[endpoint]; isAllowedMethod {
        return true
      } else {
        return false
      }
    } else if allowedMethodSet := rule[method]; (allowedMethodSet != nil) {
      if isAllowedMethod := allowedMethodSet["*"]; isAllowedMethod {
        return true
      } else if isAllowedMethod := allowedMethodSet[endpoint]; isAllowedMethod {
        return true
      } else {
        return false
      }
    } else {
      return false
    }
  }
  
  if methodFirst {
    return isAllowedMethodAndAllowedEndpoint
  } else {
    return isAllowedEndpointAndAllowedMethod
  }
}

