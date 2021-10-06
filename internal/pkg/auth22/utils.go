package auth22

import (
  "time"
  "github.com/golang-jwt/jwt/v4"
 // "errors"
)


const (
  ndaysPerYear = 365
  nhoursPerDay = 24
)


//実行されたシステム日時よりndays日後のUnix時間を取得
func CalcFutureUnixTime(ndays int) int64 {
    expiration := time.Now()
    nhours := time.Duration(ndays * nhoursPerDay) //単位は10^-9秒
    expiration = expiration.Add(nhours * time.Hour)
    return expiration.Unix()
}

//トークンの作成
func NewToken(claims map[string]interface{}, secret []byte) (string, error) {
  var jwtClaims jwt.MapClaims
  jwtClaims = claims
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
  tokenString, err := token.SignedString(secret)
  return tokenString, err 
}

//トークンの改竄検証
func ValidateSignature(tokenString string, secret []byte) (bool, error) {
  token, _ := jwt.Parse(tokenString, MakeTokenSecretSender(secret))

  if verificationTokenString, err := token.SignedString(secret); err != nil {
    return false, err
  } else {
    isSame := (tokenString == verificationTokenString)
    return isSame, nil
  }

}
//秘密鍵を送信するクロージャ
//jwtパッケージはRSA鍵ファイルの読込みなどを想定しパーサに関数を渡す仕様
func MakeTokenSecretSender(secret[]byte) func(token *jwt.Token) (interface{}, error) {
  secretSender := func(token *jwt.Token) (interface{}, error) {
    return secret, nil
  }
  return secretSender
}



