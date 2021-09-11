package user

import (
  "regexp"
  "github.com/go-ozzo/ozzo-validation/v4"
  "github.com/go-ozzo/ozzo-validation/v4/is"
  "optim_22_app/pkg/authentication"
  "optim_22_app/pkg/log"
  "optim_22_app/typefile"
  "context"
)

//#region 登録情報
//`POST /api/user`が要求する情報
type registrationInformation struct {
  Name     string `json:"name"`
  Email    string `json:"email"`
  Password string `json:"password"`
}


func (m registrationInformation) Validate() error {
  return validation.ValidateStruct(&m,
    validation.Field(&m.Name, validation.Required, validation.Length(3, 128)),
    validation.Field(&m.Email, validation.Required, is.Email),
    //is SHA256
    validation.Field(&m.Password, validation.Required, validation.Match(regexp.MustCompile("[A-Fa-f0-9]{64}$"))),
  )
}
//#endregion

type Service interface {
  Create(ctx context.Context, input registrationInformation) (string, string, error)
  Delete(ctx context.Context, userId int) error
}


type service struct {
  auth   *authentication.Authorizer
  repo   Repository
  logger log.Logger
}


func (s service) Create(ctx context.Context, req registrationInformation) (string, string, error) {
  //リクエストの値を検証
  if err := req.Validate(); err != nil {
    return "", "", err
  }
  //クエリの値を定義
  insertValues := typefile.User{
    ID:        0,         //無視される、nil `gorm:"primaryKey;autoIncrement:true"`
    Name:      req.Name,
    Email:     req.Email,
    Password:  req.Password,
  }
  //INSERTと割り当てられるuserIDを取得
  var userId int
  if err := s.repo.Create(ctx, &insertValues); err != nil {
    return "", "", err
  } else {
    userId = insertValues.ID
  }
  //トークン発行
  refreshToken, accessToken := s.auth.AuthorizationService.New(userId)

  return refreshToken, accessToken, nil
}


func (s service) Delete(ctx context.Context, userId int) error {
  //該当useriDのエントリを削除
  if err := s.repo.Delete(ctx, userId); err != nil {
    return err
  } else {
    return nil
  }
}


func StubNewService(args ...interface{}) Service { return service{nil, nil, nil}}
func StubCreate(args ...interface{}) (string, string, error)  {return "", "", nil}
func StubDelete(args ...interface{}) error {return nil}
func StubLogin(args ...interface{}) (string, string, error) {return "", "", nil}
