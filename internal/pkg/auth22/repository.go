package auth22

import (
  "context"
  "gorm.io/gorm"
  "optim_22_app/typefile"
  "optim_22_app/pkg/log"
)


type UserId struct{
  ID             int
}

type Repository interface {
  GetUserIdByCredential(ctx context.Context, credential *typefile.User) (int, error)
}


type repository struct {
  db     *gorm.DB
  logger log.Logger
}


func NewRepository(db *gorm.DB, logger log.Logger) Repository {
  return repository{db, logger}
}


func (r repository) GetUserIdByCredential(ctx context.Context, credential *typefile.User) (int, error) {
  var userId UserId
  result := r.db.WithContext(ctx).
    Model(&typefile.User{}).
    Select("id").
    Where(&credential).
    Scan(&userId)

  if err := result.Error; err != nil {
  	r.logger.Error(err)
  	return 0, err
  } else {
    r.logger.Debug(userId.ID)
  	return userId.ID, nil
  }
}