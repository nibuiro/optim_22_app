package comment

import (
  "time"
  "strconv"
  "context"
  "github.com/go-ozzo/ozzo-validation/v4"
//  "github.com/go-ozzo/ozzo-validation/v4/is"
  "optim_22_app/pkg/log"
  "optim_22_app/typefile"
//  "reflect"
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
  //is.Int  validation.Date などozzo-validationの評価関数が受け取るのは文字列のみ
  return validation.ValidateStruct(&m,
    //validation.Field(&m.UserID, validation.Required, is.Int),
    //validation.Field(&m.RequestID, validation.Required, is.Int),
    //validation.Field(&m.Date, validation.Required, validation.Date("2006-01-02")),
    validation.Field(&m.Title, validation.Required, validation.Length(3, 640)),
    validation.Field(&m.Body, validation.Required, validation.Length(3, 1280)),
    //validation.Field(&m.ReplyID, validation.Required, is.Int),
    //validation.Field(&m.Attachment, validation.Length(3, 128)),
  )
}
//#endregion


type Service interface {
  Get(ctx context.Context, req string) ([]comment, error)
  Post(ctx context.Context, req comment, requestID string) error
  Patch(ctx context.Context, req comment, requestID string) error
  Delete(ctx context.Context, requestID string, commentID string) error
}


type service struct {
  repo   Repository
  logger log.Logger
}


func NewService(repo Repository, logger log.Logger) Service {
  return service{repo, logger}
}


func (s service) Get(ctx context.Context, req string) ([]comment, error) {
  //リクエスト文字列を数値型ユーザIDに変換
  //var userId int
  requestIDAsInt, err := strconv.Atoi(req)
  if err != nil {
    return make([]comment, 1), err
  }
  //該当リクエストのコメントを取得
  if pastComments, err := s.repo.Get(ctx, requestIDAsInt); err != nil {
    return make([]comment, 1), err
  } else {
    comments := pastComments
    //nCommens := len(pastComments)
    //comments := make([]comment, pastComments)
    //for i := 0; i < nCommens; i++ {
    //  comments[i] = comment{
    //    Id: pastComments[i].ID
    //    RequestID: pastComments[i].RequestID
    //    UserID: pastComments[i].UserID
    //    UserName: pastComments[i].UserName
    //    Date: pastComments[i].Date
    //    Title: pastComments[i].Title
    //    Body: pastComments[i].Body
    //    ReplyID: pastComments[i].ReplyID
    //  }
    //}
    return comments, nil
  }
}


func (s service) Post(ctx context.Context, req comment, requestID string) error {
  //コメント登録情報を検証
  if err := req.Validate(); err != nil {
    return err
  }
  //リクエストID文字列をintに変換
  requestIDAsInt, err := strconv.Atoi(requestID)
  if err != nil {
    return err
  }
  //クエリの値を定義
  insertValues := typefile.Comment{
    RequestID: requestIDAsInt,
    UserID:    req.UserID,
    Date:      req.Date,
    Title:     req.Title,
    Body:      req.Body,
    ReplyID:   req.ReplyID,
  }

  if err := s.repo.Create(ctx, &insertValues); err != nil {
    return err
  } else {
    return nil
  }
}


func (s service) Patch(ctx context.Context, req comment, requestID string) error {
  //コメント登録情報を検証
  if err := req.Validate(); err != nil {
    return err
  }
  //リクエストID文字列を整数型に変換
  requestIDAsInt, err := strconv.Atoi(requestID)
  if err != nil {
    return err
  }
  //クエリの値を定義
  insertValues := typefile.Comment{
    RequestID: requestIDAsInt,
    UserID:    req.UserID,
    Date:      req.Date,
    Title:     req.Title,
    Body:      req.Body,
    ReplyID:   req.ReplyID,
  }

  if err := s.repo.Update(ctx, &insertValues); err != nil {
    return err
  } else {
    return nil
  }
}


func (s service) Delete(ctx context.Context, requestID string, commentID string) error{
  //コメントID文字列を整数型に変換
  commentIDAsInt, err := strconv.Atoi(commentID)
  if err != nil {
    return err
  } else {
    //コメント削除
    if err := s.repo.Delete(ctx, commentIDAsInt); err != nil {
      return err
    } else {
      return nil
    }
  }
}