package comment


import (
  "context"
  "optim_22_app/typefile"
)


func StubNewRepository(args ...interface{}) Repository {return repository{nil, nil}}


func (r repository) Get(ctx context.Context, requestID int) (typefile.Comment, error) {
  return typefile.Comment{}, nil
}


func (r repository) Create(ctx context.Context, comment *typefile.Comment) error {
  return nil
}


func (r repository) Update(ctx context.Context, comment *typefile.Comment) error {
  return nil
}


func (r repository) Delete(ctx context.Context, commentID int) error {
  return nil
}

func (r repository) DeleteByRequestID(ctx context.Context, requestID int) error {
  return nil
}