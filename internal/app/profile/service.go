package profile

import (
  "optim_22_app/pkg/log"
//  "optim_22_app/typefile"
  "encoding/json"
//  "strconv"
  "context"
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
  Get(ctx context.Context, userId string) (profile, error)
  Post(ctx context.Context, userProfile profile) error
  Patch(ctx context.Context, userProfile profile) error
  Delete(ctx context.Context, userId string) error
}


type service struct {
  repo   Repository
  logger log.Logger
}


func (s service) Get(ctx context.Context, userId string) (profile, error) {
  dummyProfile := profile{
    Bio: "test", 
    Sns: []byte(`{"twitter": "twitter.com/pole", "facebook": "facebook.com/pole"}`), 
    Submission: "test", 
    Request: "test", 
    Icon: "test",
  }
  return dummyProfile, nil
}


func (s service) Post(ctx context.Context, userProfile profile) error {
  return nil
}


func (s service) Patch(ctx context.Context, userProfile profile) error {
  return nil
}


func (s service) Delete(ctx context.Context, userId string) error {
  return nil
}


func StubNewService(args ...interface{}) Service { return service{nil, nil}}