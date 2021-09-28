package auth22


import (
  "optim_22_app/pkg/authentication"
)


func GetRule() authentication.Rule {
  return authentication.Rule{
    "POST": map[string]bool{
      "*": true,
    },
  }
}
