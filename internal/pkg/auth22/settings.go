package auth22


func GetRule() Rule {
  return Rule{
    "GET": map[string]bool{
      "*": true,
    },
    "POST": map[string]bool{
      "/auth": true,
      "/auth/refresh_token": true,
    },
  }
}
