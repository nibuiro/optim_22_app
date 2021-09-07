package test

import (
  "bytes"
  "github.com/gin-gonic/gin"
  "github.com/stretchr/testify/assert"
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
)


// APITestCase represents the data needed to describe an API test case.
type APITestCase struct {
  Name         string
  Method, URL  string
  Header       http.Header
  Cookie       *http.Cookie
  Body         string
  WantStatus   int
  WantHeader   map[string]string
  WantCookie   map[string]string
  WantBody     string
}

// Endpoint tests an HTTP endpoint using the given APITestCase spec.
func Endpoint(t *testing.T, router *gin.Engine, tc APITestCase) {
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

    stringEq(t, res.Body.String(), tc.WantResponse)

    if tc.WantHeader != nil {
      for k, v := range tc.WantHeader {
        stringEq(t, req.Header.Get(k), v)
      }
    }

    if tc.WantCookie != nil {
      for k, v := range tc.WantCookie {
        cookie, _ := req.Cookie(k)
        stringEq(t, cookie.Value, v)
      }
    }
  })
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