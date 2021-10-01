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


func NewRepository(db *gorm.DB, logger log.Logger) Repository {
  return repository{db, logger}
}


func (r repository) Create(ctx context.Context, user *typefile.User) error {
  result := r.db.WithContext(ctx).Create(user)
  if err := result.Error; err != nil {
    return err
  } else {
    //pass
  }

  client := &typefile.Client{}
  client.User.ID: user.ID
  client.User.Name: user.Name
  client.User.Email: user.Email
  client.User.Password: user.Password
  
  result := r.db.WithContext(ctx).Create(client)
  if err := result.Error; err != nil {
    return err
  } else {
    //pass
  }

  engineer := &typefile.Engineer{}
  engineer.User.ID: user.ID
  engineer.User.Name: user.Name
  engineer.User.Email: user.Email
  engineer.User.Password: user.Password
  
  result := r.db.WithContext(ctx).Create(engineer)
  if err := result.Error; err != nil {
    return err
  } else {
    //pass
  }

  return nil
}


func (r repository) Delete(ctx context.Context, userId int) error {
  result := r.db.WithContext(ctx).Delete(&typefile.User{}, userId)
  return result.Error
}


func StubNewRepository(args ...interface{}) Repository {return repository{nil, nil}}