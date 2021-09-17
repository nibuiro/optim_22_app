package profile

import (
  "optim_22_app/pkg/log"
//  "optim_22_app/typefile"
  "encoding/json"
  "strconv"
  "context"
  "errors"
)


type Sns struct {
  Twitter         string          `json:"twitter"`
  Facebook        string          `json:"facebook"`
}


type profile struct {
  Bio        string          `json:"bio"`
  Sns        json.RawMessage `json:"sns"`
  Submission string          `json:"submission"`
  Request    string          `json:"request"`
  Icon       string          `json:"icon"`
}


type Service interface {
  Get(ctx context.Context, req string) (profile, error)
  Post(ctx context.Context, req profile) error
  Patch(ctx context.Context, req profile) error
  Delete(ctx context.Context, req string) error
}


type service struct {
  repo   Repository
  logger log.Logger
}

//新たなプロフィール操作サービスを作成
func NewService(repo Repository, logger log.Logger) Service {
  return service{repo, logger}
}


func (s service) Get(ctx context.Context, req string) (profile, error) {
  //リクエスト文字列を数値型ユーザIDに変換
  var userId 
  userId, err := strconv.Atoi(req)
  if err != nil {
    return profile{}, err
  }
  //該当ユーザのプロフィールを取得
  if userProfile, err := s.repo.get(ctx, userId); err != nil {
    return profile{}, err
  } else {
    return userProfile, nil
  }
}


func (s service) Post(ctx context.Context, req profile) error {
  //リクエストの値を検証
  if err := req.Validate(); err != nil {
    return 0, err
  }
  //クエリの値を定義
  insertValues := typefile.Profile{
    Name:      req.Name,
    Email:     req.Email,
    Password:  req.Password,
  }
  //INSERT
  if err := s.repo.Create(ctx, &insertValues); err != nil {
    return err
  } else {
    return nil
  }
}


func (s service) Patch(ctx context.Context, req profile) error {
  return nil
}


func (s service) Delete(ctx context.Context, req string) error {
  return nil
}



//#region スタブ
type ServiceStub interface {
  Get(ctx context.Context, req string) (profile, error)
  Post(ctx context.Context, req profile) error
  Patch(ctx context.Context, req profile) error
  Delete(ctx context.Context, req string) error
}


type serviceStub struct {
  repo   Repository
  logger log.Logger
}


func (s serviceStub) Get(ctx context.Context, req string) (profile, error) {
  if "" == req {
    return profile{}, errors.New("不明なユーザのプロフィールを参照しました。")
  }
  dummyProfile := profile{
    Bio: "test", 
    Sns: []byte(`{"twitter": "twitter.com/pole", "facebook": "facebook.com/pole"}`), 
    Submission: "test", 
    Request: "test", 
    Icon: "test",
  }
  return dummyProfile, nil
}


func (s serviceStub) Post(ctx context.Context, req profile) error {
  return nil
}


func (s serviceStub) Patch(ctx context.Context, req profile) error {
  return nil
}


func (s serviceStub) Delete(ctx context.Context, req string) error {
  return nil
}


func NewServiceStub(args ...interface{}) ServiceStub { 
  return serviceStub{nil, nil}
}

//#endregion