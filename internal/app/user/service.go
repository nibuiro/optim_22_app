package user

import (
  "optim_22_app/pkg/log"
  "optim_22_app/typefile"
  "context"
)

//`POST /api/user`が要求する情報
type registrationInformation struct {
  name     string `json:"name"`
  email    string `json:"email"`
  password string `json:"password"`
}


//`POST /api/session`が要求する情報
type loginInformation struct {
  email    string `json:"email"`
  password string `json:"password"`
}


type Service interface {
  Create(ctx context.Context, input registrationInformation) (string, string, error)
  Delete(ctx context.Context, userId int) error
  Login(ctx context.Context, input loginInformation) (string, string, error)
}

type service struct {
  repo   Repository
  logger log.Logger
}


func (s service) Create(ctx context.Context, input registrationInformation) (string, string, error) {
  return StubCreate()
}
func (s service) Delete(ctx context.Context, userId int) error {
  return StubDelete()
}
func (s service) Login(ctx context.Context, input loginInformation) (string, string, error) {
  return StubLogin()
}




type User struct {
  typefile.User
}

type CreateUserRequest struct {
  typefile.User
}


func StubCreate(args ...interface{}) (string, string, error)  {return "", "", nil}
func StubDelete(args ...interface{}) error {return nil}
func StubLogin(args ...interface{}) (string, string, error) {return "", "", nil}



func StubNewService(args ...interface{}) Service {
  return service{nil, nil}
}
