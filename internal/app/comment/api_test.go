package comment

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

func TestGetComments(t *testing.T) {

  logger := log.New()
  router := gin.New()
  cfg, _ := config.Load("/go/src/configs/app.yaml", logger)

  repo := StubNewRepository()
  RegisterHandlers(router.Group(""), cfg, NewServiceStub(repo), logger)
  
  tests := []test.APITestCase{
    {
      Name: "dynamic url error", 
      Method: "GET", 
      URL: "/api/discussion/", 
      Header: nil, 
      Body: "",
      WantStatus: http.StatusBadRequest, 
      WantResponse: "",
    }, 
    {
      Name: "dynamic url error - no integer", 
      Method: "GET", 
      URL: "/api/discussion/hello", 
      Header: nil, 
      Body: "",
      WantStatus: http.StatusBadRequest, 
      WantResponse: "",
    }, 
    {
      Name: "Not Found", 
      Method: "GET", 
      URL: "/api/discussion/0", 
      Header: nil, 
      Body: "",
      WantStatus: http.StatusNotFound, 
      WantResponse: "",
    }, 
    {
      Name: "Json Marshal read only 1 comment", 
      Method: "GET", 
      URL: "/api/discussion/1", 
      Header: nil, 
      Body: "",
      WantStatus: http.StatusOK, 
      WantResponse: `[{"id":1,"requestID":1,"userID":1,"userName":"テスト一郎","date":"0001-01-01T00:00:00Z","title":"test","body":"test","replyID":0,"attachment":null}]`,
    },
    {
      Name: "Json Marshal read 2 comments", 
      Method: "GET", 
      URL: "/api/discussion/3", 
      Header: nil, 
      Body: "",
      WantStatus: http.StatusOK, 
      WantResponse: `[{"id":1,"requestID":1,"userID":1,"userName":"テスト一郎","date":"0001-01-01T00:00:00Z","title":"test","body":"test","replyID":0,"attachment":null},{"id":2,"requestID":1,"userID":3,"userName":"テスト三郎","date":"0001-01-01T00:00:00Z","title":"test","body":"test","replyID":0,"attachment":null}]`,
    },
  }
  for _, tc := range tests {
    test.Endpoint(t, router, tc)
  }
}



func TestPostComment(t *testing.T) {

  logger := log.New()
  router := gin.New()
  cfg, _ := config.Load("/go/src/configs/app.yaml", logger)

  repo := StubNewRepository()
  RegisterHandlers(router.Group(""), cfg, NewServiceStub(repo), logger)
  
  tests := []test.APITestCase{
    {
      Name: "Bad json parsing test", 
      Method: "POST", 
      URL: "/api/discussion/1", 
      Header: nil, 
      Body: `{"userID":1, requestID":1, "date":"2009-11-12 21:00:57", "title":"test", "body":"test", "replyID":1}`,
      WantStatus: http.StatusBadRequest, 
      WantResponse: "",
    },
  }
  for _, tc := range tests {
    test.Endpoint(t, router, tc)
  }
}


func TestDeleteComment(t *testing.T) {

  logger := log.New()
  router := gin.New()
  cfg, _ := config.Load("/go/src/configs/app.yaml", logger)

  repo := StubNewRepository()
  RegisterHandlers(router.Group(""), cfg, NewServiceStub(repo), logger)
  
  tests := []test.APITestCase{
    {
      Name: "double check test (/<invalid>/#m)", 
      Method: "DELETE", 
      URL: "/api/discussion/test/1", 
      Header: nil, 
      Body: ``,
      WantStatus: http.StatusBadRequest, 
      WantResponse: "",
    },
    {
      Name: "double check test (/#n/#m)", 
      Method: "DELETE", 
      URL: "/api/discussion/1/1", 
      Header: nil, 
      Body: ``,
      WantStatus: http.StatusOK, 
      WantResponse: "",
    },
  }
  for _, tc := range tests {
    test.Endpoint(t, router, tc)
  }
}
