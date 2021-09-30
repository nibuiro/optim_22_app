package auth22

import (
  "time"
  "regexp"
  "github.com/go-ozzo/ozzo-validation/v4"
  "github.com/golang-jwt/jwt/v4"
  "optim_22_app/typefile"
  "optim_22_app/pkg/log"
  "context"
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
  ReadCredential(data []byte) error
  ValidateCredential() error
  GenerateAccessToken() (string, error)
  GenerateRefreshToken() (string, error)
}


type service struct {
  authentication.service
  credential Credential
  repo   Repository
  logger log.Logger
}


func NewService(config *config.Config, repo Repository, logger log.Logger) Service {
  return service{
    claims: make(jwt.MapClaims)
    refreshTokenSecret: []byte(config.refreshTokenSecret)
    accessTokenSecret: []byte(config.accessTokenSecret)
    refreshTokenExpiration: config.refreshTokenExpiration
    accessTokenExpiration: config.accessTokenExpiration
    repo: repo
    logger: logger
  }
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
  if userId, err := s.repo.GetUserIdByCredential(ctx, &filter); err != nil {
    s.logger.Error(err)
    return err
  } else {
    s.claims["userID"] = userId
    return nil
  }
}


func (s service) GenerateRefreshToken() (string, error) {
  //リフレッシュトークンの期限を設定
  s.claims["exp"] = CalcFutureUnixTime(s.refreshTokenExpiration)
  //リフレッシュトークンを生成
  refreshToken, err := authentication.NewToken(s.claims["exp"], s.refreshTokenSecret)
  if err != nil {
    s.logger.Error(err)
    return "", err
  } else {
    return refreshToken, nil
  }
}


func (s service) GenerateAccessToken() (string, error) {
  //アクセストークンの期限を設定
  s.claims["exp"] = CalcFutureUnixTime(s.accessTokenExpiration)
  //アクセストークンを生成
  accessToken, err := authentication.NewToken(s.claims["exp"], s.accessTokenSecret)
  if err != nil {
    s.logger.Error(err)
    return "", err
  } else {
    return accessToken, nil
  }
}

