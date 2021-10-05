package auth22


func GetRule() Rule {
  return Rule{
    "GET": map[string]bool{
      "*": true,
    },
    "POST": map[string]bool{
      "*": true,
      "/auth": true,
      "/auth/refresh_token": true,
    },
  }
}
