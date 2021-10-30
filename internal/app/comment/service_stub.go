package comment

import (
  "context"
  "errors"
  "optim_22_app/pkg/log"
)


func NewServiceStub(args ...interface{}) Service { return serviceStub{nil, nil}}


type serviceStub struct {
  repo   Repository
  logger log.Logger
}


func (s serviceStub) Get(ctx context.Context, req string) ([]comment, error) {

  if "0" == req {
    res := make([]comment, 1)
    res[0] = comment{}
    return res, errors.New("不明なユーザのプロフィールを参照しました。")
  }

  if "1" == req {
    res := make([]comment, 1)    
    res[0] = comment{
      Id: 1,
      RequestID: 1,
      UserID: 1,
      Name: "テスト一郎",
      Title: "test",
      Body: "test",
      ReplyID: 0,
    }
    return res, nil
  }

  if "3" == req {
    res := make([]comment, 2)
    res[0] = comment{
      Id: 1,
      RequestID: 1,
      UserID: 1,
      Name: "テスト一郎",
      Title: "test",
      Body: "test",
      ReplyID: 0,
    } 
    res[1] = comment{
      Id: 2,
      RequestID: 1,
      UserID: 3,
      Name: "テスト三郎",
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


func (s serviceStub) Post(ctx context.Context, req comment, requestID string) error {
  return nil
}
