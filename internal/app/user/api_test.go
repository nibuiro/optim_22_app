package user

import (
  "net/http"
  "testing"
  "github.com/gin-gonic/gin"
  "optim_22_app/internal/pkg/test"
  "optim_22_app/pkg/log"
)
  

func TestAPI(t *testing.T) {

  logger := log.New()
  router := gin.New()
  repo := StubNewRepository()
  RegisterHandlers(router.Group(""), StubNewService(repo), logger) //auth.MockAuthHandler
  header := auth.MockAuthHeader()

  
  tests := []test.APITestCase{
    //前提を正常なbodyとして認証についてテスト
    {"create ok", "POST", "/user", `{"userID":"test", "name":"test", "email":"test", "password":"test"}`, nil, http.StatusCreated, ""},
    {"create auth error", "POST", "/user", `{"userID":"test", "name":"test", "email":"test", "password":"test"}`, header, http.StatusUnauthorized, ""},

//    bodyの検証はserviceの仕事
//    {"create missing bracket error", "POST", "/user", `"userID":"test", "name":"test", "email":"test", "password":"test"}`, nil, http.StatusBadRequest, ""},
//    {"create missing key error", "POST", "/user", `"userID":"test", "name":"test", "email":"test"}`, nil, http.StatusBadRequest, ""},
//    {"create missing value error", "POST", "/user", `"userID":"test", "name":"test", "email":"test", "password":""}`, nil, http.StatusBadRequest, ""},
//    ...
  }
  for _, tc := range tests {
    test.Endpoint(t, router, tc)
  }
}


