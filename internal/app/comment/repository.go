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


func NewRepository(db *gorm.DB, logger log.Logger) Repository {
  return repository{db, logger}
}


func (r repository) Get(ctx context.Context, requestID int) ([]comment, error) {
  //sample: SELECT comments.ID, comments.request_id, comments.user_id, users.name, comments.created_at, comments.Title, comments.Body, comments.reply_id, profiles.icon FROM comments INNER JOIN users ON comments.user_id = users.id INNER JOIN profiles ON profiles.id = users.id WHERE comments.request_id = 1
  var comments []comment
  result := r.db.WithContext(ctx).
    Model(&typefile.Comment{}).
    Select("comments.id, comments.request_id, comments.user_id, users.name, comments.created_at, comments.title, comments.body, comments.reply_id, profiles.icon").
    Joins("INNER JOIN users ON comments.user_id = users.id").
    Joins("INNER JOIN profiles ON profiles.id = users.id").
    Where("comments.request_id = ?", requestID).
    Scan(&comments)

  //r.logger.Debug(comments)

  return comments, result.Error
}


func (r repository) Create(ctx context.Context, comments *typefile.Comment) error {
  result := r.db.WithContext(ctx).Create(comments)
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


func (r repository) DeleteByRequestID(ctx context.Context, requestID int) error {
  result := r.db.WithContext(ctx).Delete(&typefile.Comment{}, requestID)
  return result.Error
}