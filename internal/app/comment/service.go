package comment

import (
  "time"
  "context"
  "github.com/go-ozzo/ozzo-validation/v4"
  "github.com/go-ozzo/ozzo-validation/v4/is"
  "optim_22_app/pkg/log"
)

//#region コメント
type comment struct {
  Id         int       `json:"id"`
  RequestID  int       `json:"requestID"`
  UserID     int       `json:"userID"`
  UserName   string    `json:"userName"`
  Date       time.Time `json:"date"`
  Title      string    `json:"title"`
  Body       string    `json:"body"`
  ReplyID    int       `json:"replyID"`
  Attachment []byte    `json:"attachment"`
}


func (m comment) Validate() error {
  return validation.ValidateStruct(&m,
    validation.Field(&m.UserID, validation.Required, is.Int),
    validation.Field(&m.RequestID, validation.Required, is.Int),
    validation.Field(&m.Date, validation.Required, validation.Date("2006-01-02")),
    validation.Field(&m.Title, validation.Required, validation.Length(3, 64)),
    validation.Field(&m.Body, validation.Required, validation.Length(3, 128)),
    validation.Field(&m.ReplyID, validation.Required, is.Int),
    //validation.Field(&m.Attachment, validation.Length(3, 128)),
  )
}
//#endregion


type Service interface {
  Get(ctx context.Context, req string) ([]comment, error)
  Post(ctx context.Context, req comment, requestID string) error
  Patch(ctx context.Context, req comment) error
  Delete(ctx context.Context, req string) error
}


type service struct {
  repo   Repository
  logger log.Logger
}

