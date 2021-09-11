package user

import (
  
  "gorm.io/gorm"
  "optim_22_app/pkg/log"
)


type Repository interface {
  Create(ctx context.Context, user *typefile.User) error
  Delete(ctx context.Context, userId string) error
}


type repository struct {
  db     *gorm.DB
  logger log.Logger
}


func (r Repository) Create(ctx context.Context, user *typefile.User) error {
  result := r.db.WithContext(ctx).Create(user)
  return result.Error
}


func StubNewRepository(args ...interface{}) Repository {
  return repository{nil}
}