package user

import (
  "optim_22_app/pkg/log"
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
  Create(ctx context.Context, input CreateUserRequest) (User, error)
}

type service struct {
  repo   Repository
  logger log.Logger
}

type User struct {
  typefile.User
}

type CreateUserRequest struct {
  typefile.User
}





func StubNewService(args ...interface{}) Service {
  return service{nil, nil}
}

