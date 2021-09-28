package auth22


import (
  "optim_22_app/pkg/authentication"
  "github.com/deckarep/golang-set"
)

const (
  Rule = authentication.Rule{
    "POST": mapset.NewSetFromSlice([]interface{"*",})
  }
)