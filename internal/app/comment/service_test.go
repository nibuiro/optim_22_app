package comment

import (
  "testing"
  "strings"
//  "github.com/gin-gonic/gin"
//  "optim_22_app/internal/pkg/test"
//  "optim_22_app/pkg/log"
//  "optim_22_app/typefile"
  "github.com/stretchr/testify/assert"
)
  

func TestTypeCommentValidate(t *testing.T) {
  /*
   *  主にutf8のカウント検証テスト
   *  utf8では日本語文字は3バイト
   */
  tests := []struct {
    name      string
    model     comment
    wantError bool
  }{
    {
      "normal query", 
      comment{
      UserID: 1,
      RequestID: 1,
      Title: strings.Repeat("あ", 210),
      Body: strings.Repeat("あ", 400),
      ReplyID: 0,
      }, 
      false,
    },
    {
      "long title", 
      comment{
      UserID: 1,
      RequestID: 1,
      Title: strings.Repeat("あ", 220),
      Body: strings.Repeat("あ", 400),
      ReplyID: 0,
      }, 
      true,
    },
    {
      "long body", 
      comment{
      UserID: 1,
      RequestID: 1,
      Title: strings.Repeat("あ", 210),
      Body: strings.Repeat("あ", 440),
      ReplyID: 0,
      }, 
      true,
    },
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      err := tt.model.Validate()
      assert.Equal(t, tt.wantError, err != nil)
    })
  }
}