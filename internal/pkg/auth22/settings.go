package auth22


import (
  "optim_22_app/pkg/authentication"
)


func GetRule() authentication.Rule {
  return authentication.Rule{
    "GET": map[string]bool{
      "*": true,
    },
  }
}
