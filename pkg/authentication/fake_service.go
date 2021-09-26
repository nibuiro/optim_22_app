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


type credential struct {
  email    string `json:"email"`
  password string `json:"password"`
}


type Service interface {
  Refresh(refreshToken string) (string, error)
  RefreshTokenSecretSender(token *jwt.Token) (interface{}, error)
  AccessTokenSecretSender(token *jwt.Token) (interface{}, error)
  ValidateCredential(ctx context.Context, req credential) (int, error)
  GenerateTokens(ctx context.Context, userID int) (string, string, error)
}


type service struct {
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


func (s service) Refresh(refreshToken string) (string, error) {
    sender :=  func (token *jwt.Token) (interface{}, error) {
      return "secret_key_for_refresh", nil
    }
    token, _ := jwt.Parse(refreshToken, sender)
    claims, _ := token.Claims.(jwt.MapClaims)

    expiration := time.Now()
    expiration = expiration.Add(time.Duration(5*365*24) * time.Hour)

    claims["exp"] = expiration.Unix()
    newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)              
    // Sign and get the complete encoded token as a string using the secret
    newTokenString, _ := newToken.SignedString([]byte("secret_key_for_refresh"))

    return newTokenString, nil
}


func (s service) ValidateCredential(ctx context.Context, req credential) (int, error) {
  //リクエストの値を検証
  return 0, nil
}


func (s service) GenerateTokens(ctx context.Context, userID int) (string, string, error) {
  //トークンを生成
  return "", "", nil
}

//パース関数にリフレッシュトークン用秘密鍵を渡すコールバック
func (s service) RefreshTokenSecretSender(token *jwt.Token) (interface{}, error) {
  return []byte("secret_key_for_refresh"), nil
}

//パース関数にアクセストークン用秘密鍵を渡すコールバック
func (s service) AccessTokenSecretSender(token *jwt.Token) (interface{}, error) {
  return []byte("secret_key"), nil
}
