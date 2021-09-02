package user

import (
  "optim_22_app/pkg/log"
)


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

