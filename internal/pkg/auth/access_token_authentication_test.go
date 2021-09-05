package auth

import (
  "net/http"
  "testing"
  "github.com/gin-gonic/gin"
  "optim_22_app/internal/pkg/test/v2"
)

const (
  refreshToken2010 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySWQiOiIwMDEiLCJleHAiOjEyODM2NTg3Mjh9.krKE34GBpQBMwSMFHf8iMpM36fxycGLvUf9Mi70--cM"
  refreshToken2020 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySWQiOiIwMDEiLCJleHAiOjE1NzgxOTMyNzl9.QrcRvgE6PbiqpAI9eLM9TeQWe6iRt0tEb-rQvnp7U_E"
  refreshToken2030 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySWQiOiIwMDEiLCJleHAiOjE5MTQ4MTA3Mjh9.QHGNRk1KMQyx8rLscYdkKxQ7nBp7ZmcLDsF8fsk40dA"
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
  * #2010: 1283658728
  * #2015: 1420426879
  * #2020: 1578193279
  * #2025: 1736046079
  * #2030: 1914810728
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

func TestRefreshTokenAuthentication(t *testing.T) {

  router := gin.New()
  logger := gin.Logger()

  //cfg.JWTExpiration => 200年 => 6307200000秒
  auth.RegisterHandlers(router.Group(""),
    auth.NewService(cfg.JWTSigningKey, 6307200000, logger),
    logger,
  )

  tests := []test.APITestCase{
    //
    {
      Name: "refresh-token refresh success", 
      Method: "POST", 
      URL: "/auth/refresh_token", 
      Header: nil, 
      Cookie: &http.Cookie{
          Name: "JWT", 
          Value: refreshToken2010, 
          HttpOnly: true, 
          Path: "/", 
          Secure: false
        },
      Body: "",
      WantStatus: http.StatusOK, 
      WantHeader: nil,
      WantCookie: map[string]string{"refresh_token": refreshToken2030},
      WantBody: ""
    },
    {
      Name: "refresh-token refresh success", 
      Method: "POST", 
      URL: "/auth/refresh_token", 
      Header: nil, 
      Cookie: &http.Cookie{
          Name: "JWT", 
          Value: refreshToken2020, 
          HttpOnly: true, 
          Path: "/", 
          Secure: false
        },
      Body: "",
      WantStatus: http.StatusOK, 
      WantHeader: nil,
      WantCookie: map[string]string{"refresh_token": refreshToken2030},
      WantBody: ""
    },
    {
      Name: "access-token refresh success [2030]", 
      Method: "POST", 
      URL: "/auth/access_token", 
      Header: nil, 
      Cookie: &http.Cookie{
          Name: "JWT", 
          Value: refreshToken2030, 
          HttpOnly: true, 
          Path: "/", 
          Secure: false
        },
      Body: "",
      WantStatus: http.StatusOK, 
      WantHeader: nil,
      WantCookie: nil,
      WantBody: `"access_token": ` + refreshToken2030
    },
    {
      Name: "access-token refresh success [2020]", 
      Method: "POST", 
      URL: "/auth/access_token", 
      Header: nil, 
      Cookie: &http.Cookie{
          Name: "JWT", 
          Value: refreshToken2020, 
          HttpOnly: true, 
          Path: "/", 
          Secure: false
        },
      Body: "",
      WantStatus: http.StatusOK, 
      WantHeader: nil,
      WantCookie: nil,
      WantBody: `"access_token": ` + refreshToken2020
    },
    {
      Name: "access-token refresh faild [2010]", 
      Method: "POST", 
      URL: "/auth/access_token", 
      Header: nil, 
      Cookie: &http.Cookie{
          Name: "JWT", 
          Value: refreshToken2010, 
          HttpOnly: true, 
          Path: "/", 
          Secure: false
        },
      Body: "",
      WantStatus: http.StatusUnauthorized, 
      WantHeader: nil,
      WantCookie: nil,
      WantBody: ""
    },
  }
  for _, tc := range tests {
    test.Endpoint(t, router, tc)
  }
}


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
      nil,
      http.StatusCreated, 
      "",
    },
    {
      "authentication failed: expired", 
      "POST", 
      "/test", 
      "", 
      MakeAuthorizationHeader(accessToken2000), 
      nil,
      http.StatusUnauthorized, 
      "",
    },
    {
      "authentication faild: no [exp] key", 
      "POST", 
      "/test", 
      "", 
      MakeAuthorizationHeader(noExpAccessToken), 
      nil,
      http.StatusUnauthorized, 
      "",
    },
    {
      "authentication faild: no [userID] key", 
      "POST", 
      "/test", 
      "", 
      MakeAuthorizationHeader(noUserIdAccessToken), 
      nil,
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
