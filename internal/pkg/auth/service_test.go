package auth

import (
  "net/http"
  "testing"
  "github.com/gin-gonic/gin"
  "optim_22_app/internal/pkg/test"
)

const (
  refreshToken2000 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEiLCJleHAiOjk2ODA0NzQxN30.mrrwDgWAPTpBK4s0PQVmRWWXWOtdGyiwMXnvZ4dfbt0"
  refreshToken2100 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEiLCJleHAiOjQxMjM3MjEwMTd9.VJbsifEaA5uaGmJdH__e270WJ20hxrlGQF79jc789vw"
  noExpRefreshToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEifQ.1LotTqA4yjwjMk9SLHPKJ3ggH2Z0j1ADVyFZqDNkZbM"
  noUserIdRefreshToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjk2ODA0NzQxN30.-pwuyxI6oFh7nkbKzdRU-3u-F6baQAMtKjwKcNRWMVo"
  accessToken2000 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEiLCJleHAiOjk2ODA0NzQxN30.x6_2QJHDmemdSz7ev6By6iyAtpWibjZLbBWZZCd3Q-U"
  accessToken2100 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEiLCJleHAiOjQxMjM3MjEwMTd9.JezlU23njkKGldV4ZH1QI37O1yCd0Y-mWmnIu-7aKEo"
  noExpAccessToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEifQ.oRHMenxs4DlJy79Has9ASiu0qD0MFh9vmYevOOksizE"
  noUserIdAccessToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjQxMjM3MjEwMTd9.pLQbblBZu_rNBqB97E6U_rFUess4Utz2tPA3aAd8mcY"
  /*
  * pythonパッケージPyJWTをによりテストトークンを生成
  *
  * #python3 -m pip install PyJWT==2.1.0
  * 
  * import jwt 
  * 
  * payload_data = {
  *     "userID": "001",
  *     "exp": 4123721017 #2100年
  * }
  * 
  * token = jwt.encode(
  *     algorithm="HS256",
  *     payload=payload_data,
  *     key='secret_key'
  * )
  * 
  * print(token)
  *
  */
)

func TestAccessTokenAuthentication(t *testing.T) {

  router := gin.New()
  logger := gin.Logger()

  RegisterMiddleware(router.Group(""), "secret_key", logger)


  router.POST("/test", func(c *gin.Context) {
    c.String(http.StatusCreated, "")
  })
  
  tests := []test.APITestCase{
    {
      "authentication success", 
      "POST", 
      "/test", 
      "", 
      MakeAuthorizationHeader(accessToken2100), 
      http.StatusCreated, 
      "",
    },
    {
      "authentication failed: expired", 
      "POST", 
      "/test", 
      "", 
      MakeAuthorizationHeader(accessToken2000), 
      http.StatusUnauthorized, 
      "",
    },
    {
      "authentication faild: no [exp] key", 
      "POST", 
      "/test", 
      "", 
      MakeAuthorizationHeader(noExpAccessToken), 
      http.StatusUnauthorized, 
      "",
    },
    {
      "authentication faild: no [userID] key", 
      "POST", 
      "/test", 
      "", 
      MakeAuthorizationHeader(noUserIdAccessToken), 
      http.StatusUnauthorized, 
      "",
    },
  }
  for _, tc := range tests {
    test.Endpoint(t, router, tc)
  }
}

func MakeAuthorizationHeader(token string) http.Header {
  header := http.Header{}
  header.Add("Authorization", token)
  return header
}
