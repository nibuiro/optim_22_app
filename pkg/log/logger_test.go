package log

import (
  "bytes"
  "context"
  "github.com/stretchr/testify/assert"
  "go.uber.org/zap"
  "net/http"
  "reflect"
  "testing"
  "os"
)

//設定ファイルが存在する場合、しない場合についてテストを実施
func TestNew(t *testing.T) {

  baseDirectory, _ := os.Getwd()

  type args struct {
    relativePath string
  }

  tests := []struct {
    name string
    args args
    want int
  }{
    //整数(0, 1)はlogger.goにてプライベート定数として宣言
    //不在であればデフォルト設定(0)
    {"no file", args{"/fake_path/fake.yaml"}, 0},
    //存在すればカスタム設定(1)
    {"zap.yaml", args{"/test/zap.yaml"}, 1},
  }

  for _, tt := range tests {

    t.Run(tt.name, func(t *testing.T) {

      absolutePath := baseDirectory + tt.args.relativePath

      l, configurationType := NewImpl(absolutePath)

      //ロガーが生成されることは設定ファイルが存在する場合、しない場合の
      //どちらでもこのテストが満たすべき必要条件
      assert.NotNil(t, l)
      //設定タイプ
      assert.Equal(t, tt.want, configurationType)

    })
  }
}

func TestNewWithZap(t *testing.T) {
  zl, _ := zap.NewProduction()
  l := NewWithZap(zl)
  assert.NotNil(t, l)
}

func TestWithRequest(t *testing.T) {
  req := buildRequest("abc", "123")
  ctx := WithRequest(context.Background(), req)
  assert.Equal(t, "abc", ctx.Value(requestIDKey).(string))
  assert.Equal(t, "123", ctx.Value(correlationIDKey).(string))

  req = buildRequest("", "123")
  ctx = WithRequest(context.Background(), req)
  assert.NotEmpty(t, ctx.Value(requestIDKey).(string))
  assert.Equal(t, "123", ctx.Value(correlationIDKey).(string))
}

func Test_getCorrelationID(t *testing.T) {
  req, _ := http.NewRequest("GET", "http://example.com", bytes.NewBufferString(""))
  assert.Empty(t, getCorrelationID(req))
  req.Header.Set("X-Correlation-ID", "test")
  assert.Equal(t, "test", getCorrelationID(req))
}

func Test_getRequestID(t *testing.T) {
  req, _ := http.NewRequest("GET", "http://example.com", bytes.NewBufferString(""))
  assert.Empty(t, getRequestID(req))
  req.Header.Set("X-Request-ID", "test")
  assert.Equal(t, "test", getRequestID(req))
}

func Test_logger_With(t *testing.T) {
  l := SilentNew()
  l2 := l.With(nil)
  assert.True(t, reflect.DeepEqual(l2, l))

  req := buildRequest("abc", "123")
  ctx := WithRequest(context.Background(), req)
  l3 := l.With(ctx)
  assert.False(t, reflect.DeepEqual(l3, l2))
}

func buildRequest(requestID, correlationID string) *http.Request {
  req, _ := http.NewRequest("GET", "http://example.com", bytes.NewBufferString(""))
  if requestID != "" {
    req.Header.Set("X-Request-ID", requestID)
  }
  if correlationID != "" {
    req.Header.Set("X-Correlation-ID", correlationID)
  }
  return req
}

func TestNewForTest(t *testing.T) {
  logger, entries := NewForTest()
  assert.Equal(t, 0, entries.Len())
  logger.Info("msg 1")
  assert.Equal(t, 1, entries.Len())
  logger.Info("msg 2")
  logger.Info("msg 3")
  assert.Equal(t, 3, entries.Len())
  entries.TakeAll()
  assert.Equal(t, 0, entries.Len())
  logger.Info("msg 4")
  assert.Equal(t, 1, entries.Len())
}