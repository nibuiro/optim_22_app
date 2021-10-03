package auth22

import (
  "context"
  "regexp"
  //"fmt"
  "github.com/go-ozzo/ozzo-validation/v4"
  "encoding/json"
  "github.com/golang-jwt/jwt/v4"
  "optim_22_app/typefile"
  "optim_22_app/pkg/log"
  "optim_22_app/internal/pkg/config"
)


type Service interface {
  ReadRefreshToken(reader jwt.MapClaims, tokenString string) (bool, error)
  ReadAccessToken(reader jwt.MapClaims, tokenString string) (bool, error)
  RefreshAccessToken(writer jwt.MapClaims) (string, error)
  RefreshRefreshToken(writer jwt.MapClaims) (string, error)
  ReadCredential(data []byte) (*Credential, error)
  ValidateCredential(ctx context.Context, reader jwt.MapClaims, writer *Credential) error
  GenerateRefreshToken(reader jwt.MapClaims) (string, error)
  GenerateAccessToken(reader jwt.MapClaims) (string, error)
}


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


type service struct {
  repo   Repository
  logger log.Logger

  RefreshTokenSecret []byte
  AccessTokenSecret []byte
  RefreshTokenExpiration int
  AccessTokenExpiration int
}


func NewService(cfg *config.Config, repo Repository, logger log.Logger) Service {
  newservice := service{
    repo: repo,
    logger: logger,

    RefreshTokenSecret: []byte(cfg.RefreshTokenSecret),
    AccessTokenSecret: []byte(cfg.AccessTokenSecret),
    RefreshTokenExpiration: cfg.RefreshTokenExpiration,
    AccessTokenExpiration: cfg.AccessTokenExpiration,
  }
//  newservice.SetParams(cfg.RefreshTokenSecret, cfg.AccessTokenSecret, cfg.RefreshTokenExpiration, cfg.AccessTokenExpiration)
  return newservice
}


func (s service) ReadRefreshToken(reader jwt.MapClaims, tokenString string) (bool, error) {
  token, err := jwt.Parse(tokenString, MakeTokenSecretSender(s.RefreshTokenSecret))
  claims, ok := token.Claims.(jwt.MapClaims)
  for key, value := range claims {
    reader[key] = value
  }
  if ok {
    return token.Valid, nil
  } else {
    return false, err
  }
}

func (s service) ReadAccessToken(reader jwt.MapClaims, tokenString string) (bool, error) {
  token, err := jwt.Parse(tokenString, MakeTokenSecretSender(s.AccessTokenSecret))
  claims, ok := token.Claims.(jwt.MapClaims)
  for key, value := range claims {
    reader[key] = value
  }
  if ok {
    return token.Valid, nil
  } else {
    return false, err
  }
}


func (s service) RefreshAccessToken(writer jwt.MapClaims) (string, error) {
    writer["exp"] = CalcFutureUnixTime(s.AccessTokenExpiration)

    newTokenString, err := NewToken(writer, s.AccessTokenSecret)
    if err != nil {
      return "", err
    } else {
      return newTokenString, nil
    }    
}


func (s service) RefreshRefreshToken(writer jwt.MapClaims) (string, error) {
    writer["exp"] = CalcFutureUnixTime(s.RefreshTokenExpiration)

    newTokenString, err := NewToken(writer, s.RefreshTokenSecret)
    if err != nil {
      return "", err
    } else {
      return newTokenString, nil
    }    
}


func (s service) ReadCredential(data []byte) (*Credential, error) {
  var reader Credential
  if err := json.Unmarshal(data, &reader); err != nil {
    s.logger.Error(err)
    return &Credential{}, err
  } else {
    return &reader, nil
  }
}


func (s service) ValidateCredential(ctx context.Context, reader jwt.MapClaims, writer *Credential) error {
  //リクエストの値を検証
  if err := writer.Validate(); err != nil {
    s.logger.Error(err)
    return err
  } 

  //idを抽出するSQL構文のWhere句の値
  filter := typefile.User{
    Email: writer.email,
    Password: writer.password,
  }
  
  //資格情報の検証とユーザIDの取
  if userId, err := s.repo.GetUserIdByCredential(ctx, &filter); err != nil {
    s.logger.Error(err)
    return err
  } else {
    reader["userID"] = userId
    return nil
  }
}


func (s service) GenerateRefreshToken(reader jwt.MapClaims) (string, error) {
  //リフレッシュトークンの期限を設定
  AddRefreshTokenExpiration(reader, s.RefreshTokenExpiration)
  //リフレッシュトークンを生成
  refreshToken, err := NewToken(reader, s.RefreshTokenSecret)
  if err != nil {
    s.logger.Error(err)
    return "", err
  } else {
    return refreshToken, nil
  }
}


func (s service) GenerateAccessToken(reader jwt.MapClaims) (string, error) {
  //アクセストークンの期限を設定
  AddAccessTokenExpiration(reader, s.AccessTokenExpiration)
  //アクセストークンを生成
  accessToken, err := NewToken(reader, s.AccessTokenSecret)
  if err != nil {
    s.logger.Error(err)
    return "", err
  } else {
    return accessToken, nil
  }
}



func AddRefreshTokenExpiration(reader jwt.MapClaims, RefreshTokenExpiration int) {
  reader["exp"] = CalcFutureUnixTime(RefreshTokenExpiration)
}


func AddAccessTokenExpiration(reader jwt.MapClaims, RefreshTokenExpiration int) {
  reader["exp"] = CalcFutureUnixTime(RefreshTokenExpiration)
}



func MakeTokenSecretSender(secret[]byte) func(token *jwt.Token) (interface{}, error) {
  secretSender := func(token *jwt.Token) (interface{}, error) {
    return secret, nil
  }
  return secretSender
}
