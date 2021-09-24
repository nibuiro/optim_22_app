package comment

import (
  "context"
  "errors"
  "time"
)


func NewServiceStub(args ...interface{}) Service { return service{nil, nil}}


func (s service) Get(ctx context.Context, req string) ([]comment, error) {

  if "0" == req {
    res := make([]comment, 1)
    res[0] = comment{}
    return res, errors.New("不明なユーザのプロフィールを参照しました。")
  }

  if "1" == req {
    res := make([]comment, 1)    
    t1, _ := time.Parse(time.RFC3339, "2016-04-13T14:12:53.4242+05:30")
    res[0] = comment{
      Id: 1,
      RequestID: 1,
      UserID: 1,
      UserName: "テスト一郎",
      Date: t1,
      Title: "test",
      Body: "test",
      ReplyID: 0,
    }
    return res, nil
  }

  if "3" == req {
    res := make([]comment, 2)
    t1, _ := time.Parse(time.RFC3339, "2016-04-13T14:12:53.4242+05:30")
    t2, _ := time.Parse(time.RFC3339, "2017-04-13T14:12:53.4242+05:30")
    res[0] = comment{
      Id: 1,
      RequestID: 1,
      UserID: 1,
      UserName: "テスト一郎",
      Date: t1,
      Title: "test",
      Body: "test",
      ReplyID: 0,
    } 
    res[1] = comment{
      Id: 2,
      RequestID: 1,
      UserID: 3,
      UserName: "テスト三郎",
      Date: t2,
      Title: "test",
      Body: "test",
      ReplyID: 0,
    }
    return res, nil
  }

  //Default
  res := make([]comment, 1)
  return res, errors.New("ハンドルされていないリクエストです")

}


func (s service) Post(ctx context.Context, req comment, requestID string) error {
  return nil
}


func (s service) Patch(ctx context.Context, req comment, requestID string) error {
  return nil
}


func (s service) Delete(ctx context.Context, requestID string, commentID string) error{
  return nil
}