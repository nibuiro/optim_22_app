package authentication

import (
  "net/http"
  "testing"
  "github.com/gin-gonic/gin"
  "time"
  "github.com/golang-jwt/jwt/v4"
  "optim_22_app/internal/pkg/test/v2"
)


//const (
//  refreshToken2010 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySWQiOiIwMDEiLCJleHAiOjEyODM2NTg3Mjh9.krKE34GBpQBMwSMFHf8iMpM36fxycGLvUf9Mi70--cM"
//  refreshToken2020 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySWQiOiIwMDEiLCJleHAiOjE1NzgxOTMyNzl9.QrcRvgE6PbiqpAI9eLM9TeQWe6iRt0tEb-rQvnp7U_E"
//  refreshToken2030 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySWQiOiIwMDEiLCJleHAiOjE5MTQ4MTA3Mjh9.QHGNRk1KMQyx8rLscYdkKxQ7nBp7ZmcLDsF8fsk40dA"
//  refreshToken2000 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEiLCJleHAiOjk2ODA0NzQxN30.mrrwDgWAPTpBK4s0PQVmRWWXWOtdGyiwMXnvZ4dfbt0"
//  refreshToken2100 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEiLCJleHAiOjQxMjM3MjEwMTd9.VJbsifEaA5uaGmJdH__e270WJ20hxrlGQF79jc789vw"
//  noExpRefreshToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEifQ.1LotTqA4yjwjMk9SLHPKJ3ggH2Z0j1ADVyFZqDNkZbM"
//  noUserIdRefreshToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjk2ODA0NzQxN30.-pwuyxI6oFh7nkbKzdRU-3u-F6baQAMtKjwKcNRWMVo"
//  accessToken2000 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEiLCJleHAiOjk2ODA0NzQxN30.x6_2QJHDmemdSz7ev6By6iyAtpWibjZLbBWZZCd3Q-U"
//  accessToken2100 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEiLCJleHAiOjQxMjM3MjEwMTd9.JezlU23njkKGldV4ZH1QI37O1yCd0Y-mWmnIu-7aKEo"
//  noExpAccessToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEifQ.oRHMenxs4DlJy79Has9ASiu0qD0MFh9vmYevOOksizE"
//  noUserIdAccessToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjQxMjM3MjEwMTd9.pLQbblBZu_rNBqB97E6U_rFUess4Utz2tPA3aAd8mcY"
//)


//期限切れのリフレッシュトークンで新しいアクセストークンは取得不可能か
func TestRefreshTokenRefreshDenied(t *testing.T) {

  router := gin.New()
  logger := gin.Logger()

  //cfg.JWTExpiration => 5年 => 157680000秒
  auth := New("secret_key_for_refresh", "secret_key", 157680000)
  //router.POST("/auth/access_token", auth.accessTokenRefreshHandler())
  router.POST("/auth/refresh_token", auth.refreshTokenRefreshHandler())
  //router.DELETE("/auth", auth.revokeHandler())

  tc := test.APITestCase{
    Name: "refresh-token refresh denied", 
    Method: "POST", 
    URL: "/auth/refresh_token", 
    Header: nil, 
    Cookie: &http.Cookie{
        Name: "JWT", 
        Value: refreshToken2020, 
        HttpOnly: true, 
        Path: "/", 
        Secure: false,
      },
    Body: "",
    WantStatus: http.StatusUnauthorized, 
    WantHeader: nil,
    WantCookie: nil,
    WantBody: "",
  }

  test.Endpoint(t, router, tc)
}


//有効期限内のリフレッシュトークンで新しいリフレッシュトークンは取得可能か
func TestRefreshTokenRefreshSuccess(t *testing.T) {

  router := gin.New()
  logger := gin.Logger()

  //cfg.JWTExpiration => 5年 => 157680000秒
  auth := New("secret_key_for_refresh", "secret_key", 157680000)
  //router.POST("/auth/access_token", auth.accessTokenRefreshHandler())
  router.POST("/auth/refresh_token", auth.refreshTokenRefreshHandler())

  tc := test.APITestCase{
    Name: "refresh-token refresh success", 
    Method: "POST", 
    URL: "/auth/refresh_token", 
    Header: nil, 
    Cookie: &http.Cookie{
        Name: "JWT", 
        Value: refreshToken2030, 
        HttpOnly: true, 
        Path: "/", 
        Secure: false,
      },
    Body: "",
    WantStatus: http.StatusOK, 
    WantHeader: nil,
    WantCookie: nil,
    WantBody: "",
  }

  t.Run(tc.Name, func(t *testing.T) {
    req, _ := http.NewRequest(tc.Method, tc.URL, bytes.NewBufferString(tc.Body))
    if tc.Header != nil {
      req.Header = tc.Header
    }
    if tc.Cookie != nil {
      req.AddCookie(tc.Cookie)
    }
    res := httptest.NewRecorder()
    if req.Header.Get("Content-Type") == "" {
      req.Header.Set("Content-Type", "application/json")
    }
    router.ServeHTTP(res, req)
    assert.Equal(t, tc.WantStatus, res.Code, "status mismatch")

    if tc.WantHeader != nil {
      for k, v := range tc.WantHeader {
        stringEq(t, req.Header.Get(k), v)
      }
    }

    //#region BodyからJWTをパースし有効期限内であることを検証
    tokenString := res.Body.String() //("access_token", _) :=

    token, _ := jwt.Parse(tokenString, accessTokenSecretSender)

    assert.True(t, token.Valid)
    //#endregion
  })
}


//期限切れのリフレッシュトークンで新しいアクセストークンは取得不可能か
func TestAccessTokenRefreshDenied(t *testing.T) {

  router := gin.New()
  logger := gin.Logger()

  //cfg.JWTExpiration => 5年 => 157680000秒
  auth := New("secret_key_for_refresh", "secret_key", 157680000)
  router.POST("/auth/access_token", auth.accessTokenRefreshHandler())
  //router.POST("/auth/refresh_token", auth.refreshTokenRefreshHandler())

  tc := test.APITestCase{
    Name: "access-token refresh denied", 
    Method: "POST", 
    URL: "/auth/access_token", 
    Header: nil, 
    Cookie: &http.Cookie{
        Name: "JWT", 
        Value: refreshToken2020, 
        HttpOnly: true, 
        Path: "/", 
        Secure: false,
      },
    Body: "",
    WantStatus: http.StatusUnauthorized, 
    WantHeader: nil,
    WantCookie: nil,
    WantBody: "",
  }

  test.Endpoint(t, router, tc)
}

//有効期限内のリフレッシュトークンで新しいアクセストークンは取得可能か
func TestAccessTokenRefreshSuccess(t *testing.T) {

  router := gin.New()
  logger := gin.Logger()

  //cfg.JWTExpiration => 5年 => 157680000秒
  auth := New("secret_key_for_refresh", "secret_key", 157680000)
  router.POST("/auth/access_token", auth.accessTokenRefreshHandler())
  //router.POST("/auth/refresh_token", auth.refreshTokenRefreshHandler())

  tc := test.APITestCase{
    Name: "access-token refresh success", 
    Method: "POST", 
    URL: "/auth/access_token", 
    Header: nil, 
    Cookie: &http.Cookie{
        Name: "JWT", 
        Value: refreshToken2030, 
        HttpOnly: true, 
        Path: "/", 
        Secure: false,
      },
    Body: "",
    WantStatus: http.StatusOK, 
    WantHeader: nil,
    WantCookie: nil,
    WantBody: "",
  }

  t.Run(tc.Name, func(t *testing.T) {
    req, _ := http.NewRequest(tc.Method, tc.URL, bytes.NewBufferString(tc.Body))
    if tc.Header != nil {
      req.Header = tc.Header
    }
    if tc.Cookie != nil {
      req.AddCookie(tc.Cookie)
    }
    res := httptest.NewRecorder()
    if req.Header.Get("Content-Type") == "" {
      req.Header.Set("Content-Type", "application/json")
    }
    router.ServeHTTP(res, req)
    assert.Equal(t, tc.WantStatus, res.Code, "status mismatch")

    if tc.WantHeader != nil {
      for k, v := range tc.WantHeader {
        stringEq(t, req.Header.Get(k), v)
      }
    }

    //#region BodyからJWTをパースし有効期限内であることを検証
    tokenString := res.Body.String() //("access_token", _) :=

    token, _ := jwt.Parse(tokenString, accessTokenSecretSender)

    assert.True(t, token.Valid)
    //#endregion
  })
}


//パース関数に秘密鍵を渡すコールバック
func accessTokenSecretSender(token *jwt.Token) (string, error) {
  return "secret_key", nil
}

//パース関数に秘密鍵を渡すコールバック
func refreshTokenSecretSender(token *jwt.Token) (string, error) {
  return "secret_key_for_refresh", nil
}

//文字列表現が等しいか確認 //JSONの表記ゆれに対応
func stringEq(t *testing.T, given string, want string) {
  if want != "" {
    pattern := strings.Trim(want, "*")
    if pattern != want {
      assert.Contains(t, given, pattern, "response mismatch")
    } else {
      assert.JSONEq(t, want, given, "response mismatch")
    }
  }
}