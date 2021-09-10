
package user
/*
import (
  "testing"
//  "github.com/gin-gonic/gin"
//  "optim_22_app/internal/pkg/test"
//  "optim_22_app/pkg/log"
//  "optim_22_app/typefile"
  "github.com/stretchr/testify/assert"
)
  

func TestCreateUserRequest_Validate(t *testing.T) {
  //バグ予測、限界値テストに基づいてテストを設計
  tests := []struct {
    name      string
    model     CreateUserRequest
    wantError bool
  }{
    {
      "success", 
      CreateUserRequest{
        Name:"test", 
        Email:"test@test.com", 
        Password:"7f83b1657ff1fc53b92dc18148a1d65dsc2d4b1fa3d677284addd200126d9069",
      }, 
      false,
    },
    {
      "success", 
      CreateUserRequest{
        Name:"test", 
        Email:"test@inc.test-ac.jp", 
        Password:"7f83b1657ff1fc53b92dc18148a1d65dsc2d4b1fa3d677284addd200126d9069",
      }, 
      false,
    },
    {
      "success", 
      CreateUserRequest{
        Name:"test", 
        Email:"test.test@test.com", 
        Password:"7f83b1657ff1fc53b92dc18148a1d65dsc2d4b1fa3d677284addd200126d9069",
      }, 
      false,
    },
    {
      "invalid name type: empty", 
      CreateUserRequest{
        Name:"", 
        Email:"test@test.com", 
        Password:"7f83b1657ff1fc53b92dc18148a1d65dsc2d4b1fa3d677284addd200126d9069",
      }, 
      true,
    },
    {
      "invalid hash type: <len", 
      CreateUserRequest{
        Name:"test", 
        Email:"test@test.com", 
        Password:"7f83b1657ff1fc53b92dc18148a1d65",
      }, 
      true,
    },
    {
      "invalid hash type: >len", 
      CreateUserRequest{
        Name:"test", 
        Email:"test@test.com", 
        Password:"7f83b1657ff1fc53b92dc18148a1d65dsc2d4b1fa3d67722d4b1fa3d677284addd200126d9069",
      }, 
      true,
    },
    {
      "invalid hash type: !format", 
      CreateUserRequest{
        Name:"test", 
        Email:"test@test.com", 
        Password:" 7f83b1657ff1fc53b92dc18148a1d65dsc2d4b1fa3d677284addd200126d9069",
      }, 
      true,
    },
    {
      "invalid hash type: !format", 
      CreateUserRequest{
        Name:"test", 
        Email:"test@test.com", 
        Password:"7f83b1657ff1fc53b92dc18148a1d65dsc2d4b1fa3d677284addd200126d9069 ",
      }, 
      true,
    },
    {
      "invalid hash type: empty", 
      CreateUserRequest{
        Name:"test", 
        Email:"test@test.com", 
        Password:"",
      }, 
      true,
    },
    {
      "invalid email type: !format", 
      CreateUserRequest{
        Name:"test", 
        Email:"testtest.com", 
        Password:"7f83b1657ff1fc53b92dc18148a1d65dsc2d4b1fa3d677284addd200126d9069",
      }, 
      true,
    },
    {
      "invalid email type: !format", 
      CreateUserRequest{
        Name:"test", 
        Email:"test@testcom", 
        Password:"7f83b1657ff1fc53b92dc18148a1d65dsc2d4b1fa3d677284addd200126d9069",
      }, 
      true,
    },
    {
      "invalid email type: empty", 
      CreateUserRequest{
        Name:"test", 
        Email:"", 
        Password:"7f83b1657ff1fc53b92dc18148a1d65dsc2d4b1fa3d677284addd200126d9069",
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
*/