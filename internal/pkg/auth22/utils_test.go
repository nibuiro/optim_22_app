package auth22

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "time"
)


func TestNewToken(t *testing.T) {
  claims := map[string]interface{}{
    "foo": "bar",
    "nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
  }
  token, err := NewToken(claims, []byte("ABC"))
  assert.Nil(t, err)
  assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.Fv2Ff2yy4AUNlZP-p0sG-y5LKwNhIOTpF7ufmDxo0yg", token)
}