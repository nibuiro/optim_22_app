package authentication


import (
  "context"
  "github.com/golang-jwt/jwt/v4"
)


type Service interface {
  WithContext(ctx context.Context) *service
  ReadRefreshToken(tokenString string) (bool, error)
  ReadAccessToken(tokenString string) (bool, error)
  RefreshTokenSecretSender(token *jwt.Token) (interface{}, error)
  AccessTokenSecretSender(token *jwt.Token) (interface{}, error)
  //オーバーライド推奨
  RefreshAccessToken() (string, error)
  RefreshRefreshToken() (string, error)
  //オーバーライド必須
  ReadCredential(data []byte) error
  ValidateCredential() error
  GenerateAccessToken() (string, error)
  GenerateRefreshToken() (string, error)
}


type service struct {
  ctx context.Context
  claims jwt.MapClaims
  refreshTokenSecret []byte
  accessTokenSecret []byte
  refreshTokenExpiration int
  accessTokenExpiration int
}

func NewService(refreshTokenSecret []byte, accessTokenSecret []byte, refreshTokenExpiration int, accessTokenExpiration int) Service {
  return service{
    refreshTokenSecret: refreshTokenSecret,
    accessTokenSecret: accessTokenSecret,
    refreshTokenExpiration: refreshTokenExpiration,
    accessTokenExpiration: accessTokenExpiration,
  }
}


func (s service) WithContext(ctx context.Context) *service {
  newServie := s
  newServie.ctx = ctx
  newServie.claims = make(jwt.MapClaims)
  return &newServie
}


func (s service) ReadRefreshToken(tokenString string) (bool, error) {
  token, err := jwt.Parse(tokenString, s.RefreshTokenSecretSender)
  claims, ok := token.Claims.(jwt.MapClaims)
  for key, value := range claims {
    s.claims[key] = value
  }
  if ok {
    return token.Valid, nil
  } else {
    return false, err
  }
}

func (s service) ReadAccessToken(tokenString string) (bool, error) {
  token, err := jwt.Parse(tokenString, s.AccessTokenSecretSender)
  claims, ok := token.Claims.(jwt.MapClaims)
  for key, value := range claims {
    s.claims[key] = value
  }
  if ok {
    return token.Valid, nil
  } else {
    return false, err
  }
}


func (s service) RefreshAccessToken() (string, error) {
    s.claims["exp"] = CalcFutureUnixTime(s.accessTokenExpiration)

    newTokenString, err := NewToken(s.claims, s.accessTokenSecret)
    if err != nil {
      return "", err
    } else {
      return newTokenString, nil
    }    
}


func (s service) RefreshRefreshToken() (string, error) {
    s.claims["exp"] = CalcFutureUnixTime(s.refreshTokenExpiration)

    newTokenString, err := NewToken(s.claims, s.refreshTokenSecret)
    if err != nil {
      return "", err
    } else {
      return newTokenString, nil
    }    
}

//パース関数にリフレッシュトークン用秘密鍵を渡すコールバック
func (s service) RefreshTokenSecretSender(token *jwt.Token) (interface{}, error) {
  return s.refreshTokenSecret, nil
}

//パース関数にアクセストークン用秘密鍵を渡すコールバック
func (s service) AccessTokenSecretSender(token *jwt.Token) (interface{}, error) {
  return s.accessTokenSecret, nil
}


func (s service) ReadCredential(data []byte) error {
  //err := json.Unmarshal(data, &s.credential)
  return nil//err
}


func (s service) ValidateCredential() error {
  //リクエストの値を検証
  return nil
}

func (s service) GenerateAccessToken() (string, error) {
  //claims := map[string]interface{}{
  //  "userID": 0,
  //}
  //トークンを生成
  return "", nil
}

func (s service) GenerateRefreshToken() (string, error) {
  //トークンを生成
  return "", nil
}


//func (s service) Refresh(refreshToken string) (string, error) {
//    sender :=  func (token *jwt.Token) (interface{}, error) {
//      return "secret_key_for_refresh", nil
//    }
//    token, _ := jwt.Parse(refreshToken, sender)
//    claims, _ := token.Claims.(jwt.MapClaims)
//
//    expiration := time.Now()
//    expiration = expiration.Add(time.Duration(5*365*24) * time.Hour)
//
//    claims["exp"] = expiration.Unix()
//    newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)              
//    // Sign and get the complete encoded token as a string using the secret
//    newTokenString, _ := newToken.SignedString([]byte("secret_key_for_refresh"))
//
//    return newTokenString, nil
//}
