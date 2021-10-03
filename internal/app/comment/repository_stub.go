package comment


import (
  "context"
  "gorm.io/gorm"
  "optim_22_app/pkg/log"
  "optim_22_app/typefile"
)


func StubNewRepository(args ...interface{}) Repository {return repositoryStub{nil, nil}}


type repositoryStub struct {
  db     *gorm.DB
  logger log.Logger
}


func (r repositoryStub) Get(ctx context.Context, requestID int) ([]comment, error) {
  return make([]comment, 1), nil
}


func (r repositoryStub) Create(ctx context.Context, comment *typefile.Comment) error {
  return nil
}


func (r repositoryStub) Update(ctx context.Context, comment *typefile.Comment) error {
  return nil
}


func (r repositoryStub) Delete(ctx context.Context, commentID int) error {
  return nil
}

func (r repositoryStub) DeleteByRequestID(ctx context.Context, requestID int) error {
  return nil
}