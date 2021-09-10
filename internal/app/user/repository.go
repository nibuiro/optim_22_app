package user

import (
  "optim_22_app/pkg/log"
)


type Repository interface {
}

type repository struct {
//  db     *dbcontext.DB
  logger log.Logger
}

func StubNewRepository(args ...interface{}) Repository {
  return repository{nil}
}