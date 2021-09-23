package comment

import (
  "context"
)


func NewServiceStub(args ...interface{}) Service { return service{nil, nil}}


func (s service) Get(ctx context.Context, req string) (Comment, error) {
  return Comment{}, nil
}


func (s service) Post(ctx context.Context, req Comment) error {
  return nil
}


func (s service) Patch(ctx context.Context, req Comment) error {
  return nil
}


func (s service) Delete(ctx context.Context, req string) error{
  return nil
}