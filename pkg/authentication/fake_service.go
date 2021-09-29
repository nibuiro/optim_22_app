package authentication


/*
 *
 *
 *  これはfake serviceです。
 *
 *
 */
import (
  "context"
  "time"
  "github.com/golang-jwt/jwt/v4"
)


type Credential struct {
  email    string `json:"email"`
  password string `json:"password"`
}


type Service interface {
  WithContext(ctx context.Context) service
  ReadRefreshToken(tokenString string) error
  ReadAccessToken(tokenString string) error
  
  ReadCredential(data []byte) error
  ValidateCredential() error
  RefreshAccessToken() (string, error)
  RefreshRefreshToken() (string, error)
  GenerateAccessToken() (string, error)
  GenerateRefreshToken() (string, error)
  RefreshTokenSecretSender() (interface{}, error)
  AccessTokenSecretSender() (interface{}, error)
}


type service struct {
  ctx context.Context
  claims jwt.MapClaims
  credential Credential
  //repo   Repository
  //logger log.Logger
}

////新たなauthenticationサービスを作成
//func NewService(repo Repository, logger log.Logger) Service {
//  return service{repo, logger}
//}


func NewService() Service {
  return service{}
}


func (s service) WithContext(ctx context.Context) service {
  newServie := s
  s.ctx = ctx
  return s
}


func (s service) ReadCredential(data []byte) error {
  err := json.Unmarshal(data, &s.credential)
  return err
}


func (s service) ReadRefreshToken(tokenString string) error {
  token, err := jwt.Parse(tokenString, s.RefreshTokenSecretSender)
  if err != nil {
    return err
  } else {
    claims, ok := token.Claims.(jwt.MapClaims)
    if ok {
      s.claims = claims
    }
  }
}

func (s service) ReadAccessToken(tokenString string) error {
  token, err := jwt.Parse(tokenString, s.AccessTokenSecretSender)
  if err != nil {
    return err
  } else {
    claims, ok := token.Claims.(jwt.MapClaims)
    if ok {
      s.claims = claims
    }
  }
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
  return "", "", nil
}

func (s service) GenerateRefreshToken() (string, error) {
  //トークンを生成
  return "", "", nil
}

//パース関数にリフレッシュトークン用秘密鍵を渡すコールバック
func (s service) RefreshTokenSecretSender() (interface{}, error) {
  return []byte("secret_key_for_refresh"), nil
}

//パース関数にアクセストークン用秘密鍵を渡すコールバック
func (s service) AccessTokenSecretSender() (interface{}, error) {
  return []byte("secret_key"), nil
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
