package profile

import (
  "context"
  "gorm.io/gorm"
  "optim_22_app/pkg/log"
  "optim_22_app/typefile"
)


type Repository interface {
  Get(ctx context.Context, userId int) (profile, error)
  Create(ctx context.Context, userProfile *profile) error
  Modify(ctx context.Context, userProfile *profile) error
  Delete(ctx context.Context, userId int) error
}


type repository struct {
  db     *gorm.DB
  logger log.Logger
}


func (s repository) Get(ctx context.Context, userId int) (profile, error) {
}


func (s repository) Create(ctx context.Context, userProfile *profile) error {
}


func (s repository) Modify(ctx context.Context, userProfile *profile) error {
}


func (s repository) Delete(ctx context.Context, userId int) error {
}


func StubNewRepository(args ...interface{}) Repository {return repository{nil, nil}}