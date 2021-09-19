package profile

import (
  "context"
  "gorm.io/gorm"
  "optim_22_app/pkg/log"
  "optim_22_app/typefile"
)


type Repository interface {
  Get(ctx context.Context, userId int) (typefile.Profile, error)
  Create(ctx context.Context, userProfile *typefile.Profile) error
  Update(ctx context.Context, userProfile *typefile.Profile) error
  Delete(ctx context.Context, userId int) error
}


type repository struct {
  db     *gorm.DB
  logger log.Logger
}


func (r repository) Get(ctx context.Context, userId int) (typefile.Profile, error) {
  var userProfile typefile.Profile
  result := r.db.WithContext(ctx).Find("ID = ?", userId, &userProfile)
  if result.Error != nil {
    return typefile.Profile{}, result.Error
  } else {
    return userProfile, nil
  }
}

func (s repository) Create(ctx context.Context, userProfile *typefile.Profile) error {
  result := r.db.WithContext(ctx).Create(userProfile)
  return result.Error
}


func (s repository) Update(ctx context.Context, userProfile *typefile.Profile) error {
  return nil
}


func (s repository) Delete(ctx context.Context, userId int) error {
  result := r.db.WithContext(ctx).Delete(&typefile.Profile{}, userId)
  return result.Error
}


func StubNewRepository(args ...interface{}) Repository {return repository{nil, nil}}