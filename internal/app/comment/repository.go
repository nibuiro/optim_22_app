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


func (r repository) Get(ctx context.Context, requestID int) ([]comment, error) {
  var comments []comment
  result := r.db.WithContext(ctx).
    Model(&typefile.Comment{}).
    Select("comment.ID, comment.RequestID, comment.UserID, user.Name, comment.Date, comment.Title, comment.Body, comment.ReplyID").
    Joins("INNER JOIN \"user\" ON comment.userID = user.ID").
    Where("comment.RequestID = ?", requestID).
    Scan(&comments)

  return comments, result.Error
}


func (r repository) Create(ctx context.Context, comment *typefile.Comment) error {
  result := r.db.WithContext(ctx).Create(comment)
  return result.Error
}


func (r repository) Update(ctx context.Context, comment *typefile.Comment) error {
  result := r.db.WithContext(ctx).Create(comment)
  return result.Error
}


func (r repository) Delete(ctx context.Context, commentID int) error {
  result := r.db.WithContext(ctx).Delete(&typefile.Comment{}, commentID)
  return result.Error
}