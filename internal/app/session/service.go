package session

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
  Create(ctx context.Context, input loginInformation) (string, string, error)
}

type service struct {
  repo   Repository
  logger log.Logger
}


func (s service) Create(ctx context.Context, input loginInformation) (string, string, error) {
  return StubCreate()
}




type User struct {
  typefile.User
}

type CreateUserRequest struct {
  typefile.User
}

func StubCreate(args ...interface{}) (string, string, error) {return "", "", nil}



func StubNewService(args ...interface{}) Service {
  return service{nil, nil}
}
