package user

import (
  
  "context"
  "gorm.io/gorm"
  "optim_22_app/typefile"
  "optim_22_app/pkg/log"
)


type Repository interface {
  Create(ctx context.Context, user *typefile.User) error
  Delete(ctx context.Context, userId int) error
}


type repository struct {
  db     *gorm.DB
  logger log.Logger
}


func (r repository) Create(ctx context.Context, user *typefile.User) error {
  result := r.db.WithContext(ctx).Create(user)
  return result.Error
}


func (r repository) Delete(ctx context.Context, userId int) error {
  result := r.db.WithContext(ctx).Delete(&typefile.User{}, userId)
  return result.Error
}


func StubNewRepository(args ...interface{}) Repository {return repository{nil, nil}}