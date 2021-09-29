package authentication


import (
  "net/http"
  "time"
  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v4"
)


type Rule map[string]map[string]bool


type resource struct {
  service Service
  domain string
  refreshTokenSecret []byte
  accessTokenSecret []byte
  refreshTokenExpiration time.Duration
  accessTokenExpiration time.Duration
}


func New(service Service, domain string, refreshTokenSecret string, accessTokenSecret string, refreshTokenExpiration time.Duration, accessTokenExpiration time.Duration) *resource {
  return &resource{
    service: service,
    domain: domain,
    refreshTokenSecret: []byte(refreshTokenSecret), 
    accessTokenSecret: []byte(accessTokenSecret), 
    refreshTokenExpiration: refreshTokenExpiration,
    accessTokenExpiration: accessTokenExpiration,
  }
}


func (rc *resource) RefreshTokenRefreshHandler() gin.HandlerFunc {
  return func(c *gin.Context) {

    if refreshToken, err := c.Cookie("refresh_token"); err != nil {
      c.Status(http.StatusBadRequest)
      return
    } else {
      token, _ := jwt.Parse(refreshToken, rc.service.RefreshTokenSecretSender)
      _, ok := token.Claims.(jwt.MapClaims)
  
      if ok {
        if token.Valid {
          newRefreshToken, _ := rc.service.Refresh(refreshToken)
          c.SetCookie("refresh_token", newRefreshToken, 1, "/", rc.domain, false, true)
            //func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
        } else {
          c.Status(http.StatusUnauthorized)
          return
        }
      } else {
        c.Status(http.StatusBadRequest)
        return
      }
    }
  }
}

//リフレッシュトークンが有効期限内のとき任意のアクセストークンの有効期限を延長
func (rc *resource) AccessTokenRefreshHandler() gin.HandlerFunc {
  return func(c *gin.Context) {

    refreshToken, err := c.Cookie("refresh_token")

    if err != nil {
      c.AbortWithStatus(http.StatusBadRequest)
    } else {

      token, _ := jwt.Parse(refreshToken, rc.service.RefreshTokenSecretSender)
      _, ok := token.Claims.(jwt.MapClaims)

      if ok {
        //リフレッシュトークンが期限内
        if token.Valid {

          //#region アクセストークンのclaimsを取り出し有効期限を更新し符号化後、署名

          //Authorizationヘッダーからstring型のトークンを取得
          tokenString := c.GetHeader("Authorization")
          //トークンの改竄と期限を検証
          token, _ := jwt.Parse(tokenString, rc.service.AccessTokenSecretSender)
          //claimsを辞書型として取得
          claims, ok := token.Claims.(jwt.MapClaims)
      
          if ok {
            expiration := time.Now()
            expiration = expiration.Add(rc.accessTokenExpiration)

            claims["exp"] = expiration.Unix()
            newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)              
            // Sign and get the complete encoded token as a string using the secret
            newTokenString, _ := newToken.SignedString([]byte(rc.accessTokenSecret))
            c.Header("Authorization", newTokenString)

            c.Status(http.StatusOK)
          } else {
            c.Status(http.StatusBadRequest)
            return
          }
          //#endregion
        } else {
          c.Status(http.StatusUnauthorized)
          return
        }
      } else {
        c.Status(http.StatusBadRequest)
        return
      }
    }
  }
}


func (rc resource) GetRefreshTokenAndAccessToken() gin.HandlerFunc {
  return func(c *gin.Context) {
  
    //BodyからJSONをパースして読み取る
    body, err := io.ReadAll(c.Request.Body)
    if err := c.BindJSON(&input); err != nil {
      c.Status(http.StatusBadRequest)
      return
    }  
    
    //資格情報の確認後、トークン生成に必要な情報を返す
    res, err := rc.service.ValidateCredential(c.Request.Context(), input)
    if err != nil {
      c.Status(http.StatusUnauthorized)
      return 
    } else {
      //資格情報確認及び認証情報取得
      refreshToken, accessToken, err := rc.service.GenerateTokens(c.Request.Context(), res)
      if err != nil {
        c.Status(http.StatusInternalServerError)
        return 
      }
      //#region ヘッダに認証情報を付加
      c.Header("Authorization", accessToken)
      c.SetCookie("refresh_token", refreshToken, 1, "/",  rc.domain, false, true)
      c.Status(http.StatusCreated)
      //#endregion
      return 
    }
  }
}


func (rc resource) RefreshAccessTokenAndRefreshToken() gin.HandlerFunc {
  return func(c *gin.Context) {

    refreshToken, err := c.Cookie("refresh_token")

    if err != nil {
      c.Status(http.StatusBadRequest)
      return
    } else {

      token, _ := jwt.Parse(refreshToken, rc.service.RefreshTokenSecretSender)
      claims, ok := token.Claims.(jwt.MapClaims)
  
      if ok {
        if token.Valid {
          //資格情報確認及び認証情報取得
          refreshToken, accessToken, err := rc.service.GenerateTokens(c.Request.Context(), claims)
          if err != nil {
            return 
          }
          //#region ヘッダに認証情報を付加
          c.Header("Authorization", accessToken)
          c.SetCookie("refresh_token", refreshToken, 1, "/",  rc.domain, false, true)
          c.Status(http.StatusCreated)
          //#endregion
          return 
        } else {
          c.Status(http.StatusUnauthorized)
          return
        }
      } else {
        c.Status(http.StatusBadRequest)
        return
      }
    }
  }
}

//認証情報を空文字列で上書き
func (rc *resource) RevokeHandler() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.SetCookie("refresh_token", "", 0, "/", "", false, true)
    c.Header("Authorization", "")
  }
}