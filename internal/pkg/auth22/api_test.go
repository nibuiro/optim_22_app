package auth22

import (
  "net/http"
)




func MakeAuthorizationHeader(refreshToken string, accessToken string) http.Header {
  header := http.Header{}
  header.Add("Authorization", accessToken)
  header.Add("Refresh-Token", refreshToken)
  return header
}

/*
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
*/
