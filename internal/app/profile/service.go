package profile

import (
  "regexp"
  "github.com/go-ozzo/ozzo-validation/v4"
  "optim_22_app/pkg/log"
  "optim_22_app/typefile"
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
  Id         string          `json:"userID"`
  Bio        string          `json:"bio"`
  Sns        json.RawMessage `json:"sns"`
  Submission json.RawMessage `json:"submission"`
  Request    json.RawMessage `json:"request"`
  Icon       string          `json:"icon"`
}


func (m profile) Validate() error {
  return validation.ValidateStruct(&m,
    validation.Field(&m.Id, validation.Required, validation.Length(3, 128)),
    //is.Email@ozzo-validation/v4/isはテストケース`success#1`にてエラー
    validation.Field(&m.Bio, validation.Required, validation.Match(regexp.MustCompile("[a-zA-Z]+[a-zA-Z0-9\\.]@[a-zA-Z]+((\\.[a-zA-Z0-9\\-])+[a-zA-Z0-9]+)+"))),
    //is SHA256
    validation.Field(&m.Sns, validation.Required, validation.Length(64, 64), validation.Match(regexp.MustCompile("[A-Fa-f0-9]{64}$"))),
    validation.Field(&m.Icon, validation.Required, validation.Length(64, 64), validation.Match(regexp.MustCompile("[A-Fa-f0-9]{64}$"))),
  )
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
  //var userId int
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
  //SNS登録情報を読み込み
  sns := Sns{}
  json.Unmarshal(req.Sns, &sns)
  //SNS登録情報を検証
  if err := sns.Validate(); err != nil {
    return err
  }
  //リクエストの値を検証
  if err := req.Validate(); err != nil {
    return err
  }
  //クエリの値を定義
  insertValues := typefile.Profile{
    ID:      req.Id,
    Bio:     req.Bio,
    Sns:     req.Sns,
    Icon:    req.Icon,
  }
  //INSERT
  if err := s.repo.Create(ctx, &insertValues); err != nil {
    return err
  } else {
    return nil
  }
}


func (s service) Patch(ctx context.Context, req profile) error {
  //SNS登録情報を読み込み
  sns := Sns{}
  json.Unmarshal(req.Sns, &sns)
  //SNS登録情報を検証
  if err := sns.Validate(); err != nil {
    return err
  }
  //リクエストの値を検証
  if err := req.Validate(); err != nil {
    return err
  }
  //クエリの値を定義
  insertValues := typefile.Profile{
    ID:      req.Id,
    Bio:     req.Bio,
    Sns:     req.Sns,
    Icon:    req.Icon,
  }
  //UPDATE
  if err := s.repo.Update(ctx, &insertValues); err != nil {
    return err
  } else {
    return nil
  }
}


func (s service) Delete(ctx context.Context, req string) error {
  //リクエスト文字列を数値型ユーザIDに変換
  //var userId int
  userId, err := strconv.Atoi(req)
  if err != nil {
    return err
  }
  //該当ユーザのプロフィールを削除
  if err := s.repo.delete(ctx, userId); err != nil {
    return err
  } else {
    return　nil
  }
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