package authentication

import (
  "net/http"
//  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v4"
)

type Ruler func(string, string) bool

func (rc *resource) ValidateAccessToken(rule Rule, methodFirst bool) gin.HandlerFunc {

  IsRestricted := makeRuler(rule, methodFirst)

  return func(c *gin.Context) {

    if IsRestricted(c.Request.Method, c.FullPath()) {
      //Authorizationヘッダーからstring型のトークンを取得
      tokenString := c.GetHeader("Authorization")
      //トークンの改竄と期限を検証
      //tips: expキーがない場合無期限トークンとして扱われ、token.Validの値はtrue
      token, _ := jwt.Parse(tokenString, rc.service.AccessTokenSecretSender)
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

  isRestrictedMethodAndRestrictedEndpoint := func (method string, endpoint string) bool {

    if restrictedEndpointSet := rule["*"]; (restrictedEndpointSet != nil) {
      if isRestrictedEndpoint := restrictedEndpointSet["*"]; isRestrictedEndpoint {
        return true
      } else if isRestrictedEndpoint := restrictedEndpointSet[endpoint]; isRestrictedEndpoint {
        return true
      } else {
        return false
      }
    } else if restrictedEndpointSet := rule[method]; (restrictedEndpointSet != nil) {
      if isRestrictedEndpoint := restrictedEndpointSet["*"]; isRestrictedEndpoint {
        return true
      } else if isRestrictedEndpoint := restrictedEndpointSet[endpoint]; isRestrictedEndpoint {
        return true
      } else {
        return false
      }
    } else {
      return false
    }
  }
  
  isRestrictedEndpointAndRestrictedMethod := func (method string, endpoint string) bool {
    if restrictedMethodSet := rule["*"]; (restrictedMethodSet != nil) {
      if isRestrictedMethod := restrictedMethodSet["*"]; isRestrictedMethod {
        return true
      } else if isRestrictedMethod := restrictedMethodSet[endpoint]; isRestrictedMethod {
        return true
      } else {
        return false
      }
    } else if restrictedMethodSet := rule[method]; (restrictedMethodSet != nil) {
      if isRestrictedMethod := restrictedMethodSet["*"]; isRestrictedMethod {
        return true
      } else if isRestrictedMethod := restrictedMethodSet[endpoint]; isRestrictedMethod {
        return true
      } else {
        return false
      }
    } else {
      return false
    }
  }
  
  if methodFirst {
    return isRestrictedMethodAndRestrictedEndpoint
  } else {
    return isRestrictedEndpointAndRestrictedMethod
  }
}

