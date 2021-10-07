package comment

import (
  "net/http"
  "testing"
  "github.com/gin-gonic/gin"
  "optim_22_app/internal/pkg/test"
  "optim_22_app/pkg/log"
)
  
/* 
 * エンドポイントとして機能するか？
 */

func TestGetComments(t *testing.T) {

  logger := log.New()
  router := gin.New()

  repo := StubNewRepository()
  RegisterHandlers(router.Group(""), NewServiceStub(repo), logger)
  
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
      WantResponse: `[{"comment_id":1,"request_id":1,"user_id":1,"username":"テスト一郎","title":"test","text":"test","createdat":"0001-01-01T00:00:00Z","reply_id":0,"attachment":null,"icon":null}]`,
    },
    {
      Name: "Json Marshal read 2 comments", 
      Method: "GET", 
      URL: "/api/discussion/3", 
      Header: nil, 
      Body: "",
      WantStatus: http.StatusOK, 
      WantResponse: `[{"comment_id":1,"request_id":1,"user_id":1,"username":"テスト一郎","title":"test","text":"test","createdat":"0001-01-01T00:00:00Z","reply_id":0,"attachment":null,"icon":null},{"comment_id":2,"request_id":1,"user_id":3,"username":"テスト三郎","title":"test","text":"test","createdat":"0001-01-01T00:00:00Z","reply_id":0,"attachment":null,"icon":null}]`,
    },
  }
  for _, tc := range tests {
    test.Endpoint(t, router, tc)
  }
}



func TestPostComment(t *testing.T) {

  logger := log.New()
  router := gin.New()

  repo := StubNewRepository()
  RegisterHandlers(router.Group(""), NewServiceStub(repo), logger)
  
  tests := []test.APITestCase{
    {
      Name: "Bad json parsing test", 
      Method: "POST", 
      URL: "/api/discussion/1", 
      Header: nil, 
      Body: `{"user_id":1, request_id":1, "date":"2016-04-13T14:12:53.4242+05:30", "title":"test", "text":"test", "reply_id":1}`,
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

  repo := StubNewRepository()
  RegisterHandlers(router.Group(""), NewServiceStub(repo), logger)
  
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
