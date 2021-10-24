package user

import (
  
  "context"
  "gorm.io/gorm"
  "optim_22_app/typefile"
  "optim_22_app/pkg/log"
)


type Repository interface {
  Create(ctx context.Context, user *typefile.User) error
}


type repository struct {
  db     *gorm.DB
  logger log.Logger
}


func NewRepository(db *gorm.DB, logger log.Logger) Repository {
  return repository{db, logger}
}


func (r repository) Create(ctx context.Context, user *typefile.User) error {
  
  tx := r.db.WithContext(ctx).Begin()
  //ユーザ登録以降基本的に編集を受け付ける形となり初期エントリが必要
  err := CreateInitialEntries(tx, user)
  defer func() {
    if err != nil {
      tx.Rollback()
    } else {
      tx.Commit()
    }
  }()

  if err != nil {
    return err
  } else {
    return nil
  }
}


func CreateInitialEntries(tx *gorm.DB, user *typefile.User) error {
  if err := tx.Create(user).Error; err != nil {
    return err
  } else {
    //pass
  }

  client := &typefile.Client{}
  client.User.ID = user.ID
  client.User.Name = user.Name
  client.User.Email = user.Email
  //client.User.Password = user.Password
  
  if err := tx.Create(client).Error; err != nil {
    return err
  } else {
    //pass
  }

  engineer := &typefile.Engineer{}
  engineer.User.ID = user.ID
  engineer.User.Name = user.Name
  engineer.User.Email = user.Email
  //engineer.User.Password = user.Password
  
  if err := tx.Create(engineer).Error; err != nil {
    return err
  } else {
    //pass
  }

  profile := &typefile.Profile{}
  profile.ID = user.ID
  profile.Bio = ``
  profile.Sns = []byte(`{"github":"","twitter":"","facebook":""}`)
  profile.Icon =  `data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAAMSURBVBhXY/j//z8ABf4C/qc1gYQAAAAASUVORK5CYII=`

  
  if err := tx.Create(profile).Error; err != nil {
    return err
  } else {
    //pass
  }

  return nil
}


func StubNewRepository(args ...interface{}) Repository {return repository{nil, nil}}