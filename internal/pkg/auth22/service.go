package auth22

import (
  "optim_22_app/typefile"
  "optim_22_app/pkg/log"
  "optim_22_app/typefile"
  "context"
  "optim_22_app/pkg/authentication"
)


//`POST /api/auth`が要求する情報
type credential struct {
  email    string `json:"email"`
  password string `json:"password"`
}


func (m credential) Validate() error {
  return validation.ValidateStruct(&m,
    //is.Email@ozzo-validation/v4/isはテストケース`success#1`にてエラー
    validation.Field(&m.Email, validation.Required, validation.Match(regexp.MustCompile("[a-zA-Z]+[a-zA-Z0-9\\.]@[a-zA-Z]+((\\.[a-zA-Z0-9\\-])+[a-zA-Z0-9]+)+"))),
    //is SHA256
    validation.Field(&m.Password, validation.Required, validation.Length(64, 64), validation.Match(regexp.MustCompile("[A-Fa-f0-9]{64}$"))),
  )
}


type Service interface {
  ValidateCredential(ctx context.Context, req credential) (int, error)
  GenerateTokens(ctx context.Context, claims map[string]interface{}) (string, string, error)
}


type service struct {
  repo   Repository
  logger log.Logger
}


func (s service) ValidateCredential(ctx context.Context, req credential) (map[string]interface{}, error) {
  //リクエストの値を検証
  if err := req.Validate(); err != nil {
    s.logger.Error(err)
    return 0, err
  }

  //idを抽出するSQL構文のWhere句の値
  filter = typefile.User{
    Email: req.email
    Password: req.password
  }
  
  //資格情報の検証とユーザIDの取得
  var userId int
  if userId, err := s.repo.GetUserIdByCredential(ctx, filter); err != nil {
    return 0, err
  } else {
    claims := map[string]interface{}{
      "userID": userId,
    }
    return claims, nil
  }
}

func (s service) GenerateTokens(ctx context.Context, claims map[string]interface{}) (string, string, error) {

  userID := claims["userID"].(int)

  claims := map[string]interface{}{
    "userid": userID,
    "exp": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
  }

  if refreshToken, err := authentication.NewToken(claims, s.config.refreshTokenSecretKey); err != nil {
    s.logger.Error(err)
    return "", "", err
  } else {
    if accessToken, err := authentication.NewToken(claims, s.config.accessTokenSecretKey); err != nil {
      s.logger.Error(err)
      return "", "", err
    } else {
      refreshToken, accessToken, nil
    }
  }
 
}

//パース関数にリフレッシュトークン用秘密鍵を渡すコールバック
func (s service) refreshTokenSecretSender(token *jwt.Token) (interface{}, error) {
  return s.config.refreshTokenSecret, nil
}

//パース関数にアクセストークン用秘密鍵を渡すコールバック
func (s service) accessTokenSecretSender(token *jwt.Token) (interface{}, error) {
  return s.config.accessTokenSecret, nil
}
