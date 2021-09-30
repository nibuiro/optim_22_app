package auth22

import (
  "context"
  "regexp"
  "github.com/go-ozzo/ozzo-validation/v4"
  "encoding/json"
  "github.com/golang-jwt/jwt/v4"
  "optim_22_app/typefile"
  "optim_22_app/pkg/log"
  "optim_22_app/pkg/authentication"
  "optim_22_app/internal/pkg/config"
)


//`POST /api/auth`が要求する情報
type Credential struct {
  email    string `json:"email"`
  password string `json:"password"`
}


func (m Credential) Validate() error {
  return validation.ValidateStruct(&m,
    //is.Email@ozzo-validation/v4/isはテストケース`success#1`にてエラー
    validation.Field(&m.email, validation.Required, validation.Match(regexp.MustCompile("[a-zA-Z]+[a-zA-Z0-9\\.]@[a-zA-Z]+((\\.[a-zA-Z0-9\\-])+[a-zA-Z0-9]+)+"))),
    //is SHA256
    validation.Field(&m.password, validation.Required, validation.Length(64, 64), validation.Match(regexp.MustCompile("[A-Fa-f0-9]{64}$"))),
  )
}


type Service interface {  
  SetParams(refreshTokenSecret string, accessTokenSecret string, refreshTokenExpiration int, accessTokenExpiration int)
  //
  SetContext(ctx context.Context)
  WithContext(ctx context.Context) *service //オーバーライド必須
  WhereContext() context.Context
  //
  RefreshClaims()
  SetClaims(key string, value interface{})
  GetClaims() jwt.MapClaims
  AddRefreshTokenExpiration()
  AddAccessTokenExpiration()
  GetSignedRefreshToken() (string, error)
  GetSignedAccessToken() (string, error)
  //
  ReadRefreshToken(tokenString string) (bool, error)
  ReadAccessToken(tokenString string) (bool, error)
  RefreshTokenSecretSender(token *jwt.Token) (interface{}, error)
  AccessTokenSecretSender(token *jwt.Token) (interface{}, error)
  //
  RefreshAccessToken() (string, error) //オーバーライド推奨
  RefreshRefreshToken() (string, error) //オーバーライド推奨
  //
  ReadCredential(data []byte) error //オーバーライド必須
  ValidateCredential() error //オーバーライド必須
  GenerateAccessToken() (string, error) //オーバーライド必須
  GenerateRefreshToken() (string, error) //オーバーライド必須
}


type service struct {
  authentication.Service
  credential Credential
  repo   Repository
  logger log.Logger
}


func NewService(config *config.Config, repo Repository, logger log.Logger) Service {
  newService := service{
    repo: repo,
    logger: logger,
  }
  newService.SetParams(config.RefreshTokenSecret, config.AccessTokenSecret, config.RefreshTokenExpiration, config.AccessTokenExpiration)
  return newService
}


func (s service) WithContext(ctx context.Context) *service {
  newServie := s
  newServie.SetContext(ctx)
  newServie.RefreshClaims()
  return &newServie
}


func (s service) ReadCredential(data []byte) error {
  if err := json.Unmarshal(data, &s.credential); err != nil {
    s.logger.Error(err)
    return err
  } else {
    return nil
  }
}


func (s service) ValidateCredential() error {
  //リクエストの値を検証
  if err := s.credential.Validate(); err != nil {
    s.logger.Error(err)
    return err
  }

  //idを抽出するSQL構文のWhere句の値
  filter := typefile.User{
    Email: s.credential.email,
    Password: s.credential.password,
  }
  
  //資格情報の検証とユーザIDの取
  if userId, err := s.repo.GetUserIdByCredential(s.WhereContext(), &filter); err != nil {
    s.logger.Error(err)
    return err
  } else {
    s.SetClaims("userID", userId)
    return nil
  }
}


func (s service) GenerateRefreshToken() (string, error) {
  //リフレッシュトークンの期限を設定
  s.AddRefreshTokenExpiration()
  //リフレッシュトークンを生成
  refreshToken, err := s.GetSignedRefreshToken()
  if err != nil {
    s.logger.Error(err)
    return "", err
  } else {
    return refreshToken, nil
  }
}


func (s service) GenerateAccessToken() (string, error) {
  //アクセストークンの期限を設定
  s.AddAccessTokenExpiration()
  //アクセストークンを生成
  accessToken, err := s.GetSignedAccessToken()
  if err != nil {
    s.logger.Error(err)
    return "", err
  } else {
    return accessToken, nil
  }
}

