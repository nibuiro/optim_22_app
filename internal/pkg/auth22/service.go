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


func NewService(ctx context.Context, config *config.Config, repo Repository, logger log.Logger) Service {
  return service{
    ctx: ctx
    claims: make(jwt.MapClaims)
    refreshTokenSecret: config.refreshTokenSecret
    accessTokenSecret: config.accessTokenSecret
    refreshTokenExpiration: config.refreshTokenExpiration
    accessTokenExpiration: config.accessTokenExpiration
    repo: repo
    logger: logger
  }
}


func (s service) ValidateCredential() (map[string]interface{}, error) {
  //リクエストの値を検証
  if err := req.Validate(); err != nil {
    s.logger.Error(err)
    return nil, err
  }

  //idを抽出するSQL構文のWhere句の値
  filter := typefile.User{
    Email: req.email,
    Password: req.password,
  }
  
  //資格情報の検証とユーザIDの取
  if userId, err := s.repo.GetUserIdByCredential(ctx, &filter); err != nil {
    return nil, err
  } else {
    claims := map[string]interface{}{
      "userID": userId,
    }
    return claims, nil
  }
}

func (s service) GenerateTokens(ctx context.Context, claims map[string]interface{}) (string, string, error) {

  //userID := claims["userID"].(int)

  expiration := time.Now()
  expiration = expiration.Add(authentication.CalcYears2SecondsConversion(s.config.RefreshTokenExpiration))

  claims["exp"] = expiration.Unix()
  //:= map[string]interface{}{
  //  "userid": userID,
  //  "exp": expiration.Unix(),
  //}

  if refreshToken, err := authentication.NewToken(claims, s.config.RefreshTokenSecretKey); err != nil {
    s.logger.Error(err)
    return "", "", err
  } else {
    if accessToken, err := authentication.NewToken(claims, s.config.AccessTokenSecretKey); err != nil {
      s.logger.Error(err)
      return "", "", err
    } else {
      return refreshToken, accessToken, nil
    }
  }
 
}

