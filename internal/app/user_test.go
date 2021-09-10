package user

import (
  "net/http"
  "testing"
  "strconv"
  "github.com/gin-gonic/gin"
  "optim_22_app/internal/pkg/test"
  "optim_22_app/pkg/log"
)
  
/* 
 * APIとして機能するか？
 */

//ユーザ作成時にIDが決定されトークンが発行されるか
func TestGetNewIdWithToken(t *testing.T) {

  logger := log.New()
  router := gin.New()
  repo := StubNewRepository()
  RegisterHandlers(router.Group(""), StubNewService(repo), logger)

  testCase := test.APITestCase{
    Name: "register success", 
    Method: "POST", 
    URL: "/api/user", 
    Header: nil, 
    Body: `{name":"test", "email":"test@test.test", "password":"test"}`,
    WantStatus: http.StatusCreated, 
    WantResponse: "",
  }
  
  //ヘッダとCookie以外について検証
  res := test.Endpoint(t, router, testCase)

  //#region ヘッダからJWTをパースし有効期限内であることを検証
  tokenString := res.HeaderMap["Authorization"][0]
  token, _ := jwt.Parse(tokenString, "secret_key")
  assert.True(t, token.Valid)
  //#endregion

  //#region CookieからJWTをパースし有効期限内であることを検証
  onlyCookiePacket := &http.Request{Header: http.Header{"Cookie": res.HeaderMap["Set-Cookie"]}}
  cookie, _  := onlyCookiePacket.Cookie("refresh_token")
  tokenString = cookie.Value
  token, _ = jwt.Parse(tokenString, auth.refreshTokenSecretSender)
  assert.True(t, token.Valid)
  //#endregion

  //#region CookieからJWTをパースし有効期限内であることを検証
  claims, _ := token.Claims.(jwt.MapClaims)
  i, _ := strconv.Atoi(claims["userID"])
  assert.Equal(t, 0, i)
  //#endregion
}

//同一メールアドレスによる多重登録の検証
func TestDoubleRegistration(t *testing.T) {

  logger := log.New()
  router := gin.New()
  repo := StubNewRepository()
  RegisterHandlers(router.Group(""), StubNewService(repo), logger)

  testCase := test.APITestCase{
    Name: "register success", 
    Method: "POST", 
    URL: "/api/user", 
    Header: nil, 
    Body: `{name":"testA", "email":"test@test.test", "password":"test"}`,
    WantStatus: http.StatusCreated, 
    WantResponse: "",
  }
  
  //ヘッダとCookie以外について検証
  res := test.Endpoint(t, router, testCase)

  //名前を変えて同一メールアドレスで再登録
  testCase := test.APITestCase{
    Name: "register success", 
    Method: "POST", 
    URL: "/api/user", 
    Header: nil, 
    Body: `{name":"testB", "email":"test@test.test", "password":"test"}`,
    WantStatus: http.StatusBadRequest, 
    WantResponse: "",
  }
  
  //ヘッダとCookie以外について検証
  res := test.Endpoint(t, router, testCase)
}


func MakeAuthorizationHeader(token string, cookies []http.Cookie) http.Header {
  header := http.Header{}
  header.Add("Authorization", token)

  cookieCount := len(cookies)
  if cookieCount != 0 {
    for _, cookie := range cookies {
      header.Add("Cookie", cookie.String())
    }
  }
  return header
}

  





