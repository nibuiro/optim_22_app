package auth22


func GetRule() Rule {
  return Rule{
    "GET": map[string]bool{
      "*": true,
    },
  }
}
