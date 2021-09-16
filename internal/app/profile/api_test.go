package profile

import (
  "net/http"
  "testing"
  "github.com/gin-gonic/gin"
  "optim_22_app/internal/pkg/test"
  "optim_22_app/pkg/log"
  "optim_22_app/internal/pkg/config"
)
  
/* 
 * エンドポイントとして機能するか？
 */

func TestPostProfile(t *testing.T) {

  logger := log.New()
  router := gin.New()
  cfg, _ := config.Load("/go/src/configs/app.yaml", logger)

  repo := StubNewRepository()
  RegisterHandlers(router.Group(""), cfg, StubNewService(repo), logger)
  
  tests := []test.APITestCase{
    {
      Name: "Nested json parsing test", 
      Method: "POST", 
      URL: "/api/profile", 
      Header: nil, 
      Body: `{"bio":"test", "sns":{"twitter": "twitter.com/pole", "facebook": "facebook.com/pole"}, "submission":"test", "request":"test", "icon":"test"}`,
      WantStatus: http.StatusCreated, 
      WantResponse: "",
    }, 
    {
      Name: "Bad json parsing test", 
      Method: "POST", 
      URL: "/api/profile", 
      Header: nil, 
      Body: `"bio":"test", "sns":{"twitter": "twitter.com/pole", "facebook": "facebook.com/pole"}, "submission":"test", "request":"test", "icon":"test"`,
      WantStatus: http.StatusBadRequest, 
      WantResponse: "",
    },
  }
  for _, tc := range tests {
    test.Endpoint(t, router, tc)
  }
}


