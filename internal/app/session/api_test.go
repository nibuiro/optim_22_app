package session

import (
  "net/http"
  "testing"
  "github.com/gin-gonic/gin"
  "optim_22_app/internal/pkg/test/v2"
  "optim_22_app/pkg/log"
  "optim_22_app/internal/pkg/config"
  "github.com/stretchr/testify/assert"
)
  
/* 
 * インターフェイスに持たせている機能についてテスト
 */



//Cookieの"refresh_token"の値が空文字で上書きする命令が発行されるか
func TestLogout(t *testing.T) {

  logger := log.New()
  router := gin.New()
  repo := StubNewRepository()
  rc := resource{&config.Config{"localhost", 8080, "dummy", "dummy"}, StubNewService(repo), logger}

  router.DELETE("/api/session", rc.delete())

  testCase := test.APITestCase{
    Name: "logout success (overwite cookie)", 
    Method: "DELETE", 
    URL: "/api/session", 
    Header: nil, 
    Body: "",
    WantStatus: http.StatusOK, 
    WantResponse: "",
  }
  
  //ヘッダとCookie以外について検証
  res := test.Endpoint(t, router, testCase)

  onlyCookiePacket := &http.Request{Header: http.Header{"Cookie": res.HeaderMap["Set-Cookie"]}}
  cookie, _  := onlyCookiePacket.Cookie("refresh_token")
  assert.Equal(t, "", cookie.Value)

}