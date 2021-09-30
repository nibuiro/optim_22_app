package authentication

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

    s := rc.service.WithContext(c.Request.Context())

    if !IsAllowed(c.Request.Method, c.FullPath()) {
      //Authorizationヘッダーからstring型のトークンを取得
      tokenString := c.GetHeader("Authorization")
      //トークンの改竄と期限を検証
      //tips: expキーがない場合無期限トークンとして扱われ、token.Validの値はtrue
      token, _ := jwt.Parse(tokenString, s.AccessTokenSecretSender)
      //辞書型に変換
      _, ok := token.Claims.(jwt.MapClaims)
      //claims, ok := token.Claims.(jwt.MapClaims)
  
      if ok {
        if token.Valid {
          //なにもしない
        } else {
          c.AbortWithStatus(http.StatusUnauthorized)
        }
      } else {
        c.AbortWithStatus(http.StatusBadRequest)
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

