package comment


import (
  "context"
  "gorm.io/gorm"
  "optim_22_app/pkg/log"
  "optim_22_app/typefile"
)


type Repository interface {
  Get(ctx context.Context, requestID int) ([]comment, error)
  Create(ctx context.Context, comment *typefile.Comment) error
  Update(ctx context.Context, comment *typefile.Comment) error
  Delete(ctx context.Context, commentID int) error

  DeleteByRequestID(ctx context.Context, requestID int) error
}


type repository struct {
  db     *gorm.DB
  logger log.Logger
}


