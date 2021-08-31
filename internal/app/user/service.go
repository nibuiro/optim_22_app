package user

import (
  "optim_22_app/pkg/log"
)


type Service interface {
}

type service struct {
  repo   Repository
  logger log.Logger
}

func StubNewService(args ...interface{}) Service {
  return service{nil, nil}
}