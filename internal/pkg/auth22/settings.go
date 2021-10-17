package auth22

//許可するURLを指定 //"*"はワイルドカード
func GetRule() Rule {
  return Rule{
    "GET": map[string]bool{
      "*": true,
    },
    "PUT": map[string]bool{
      //"*": true,
    },
    "POST": map[string]bool{
      //"*": true,
      "/auth": true,
      "/auth/refresh_token": true,
      "/api/user": true, //ユーザ登録エンドポイントを露出
    },
  }
}
