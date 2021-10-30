package auth22

/*
 * * [jwt/hmac_example_test.go · golang-jwt/jwt](https://github.com/golang-jwt/jwt/blob/2ebb50f957d606de5909fcf9ed49f9af3bc35e97/hmac_example_test.go#L54)
 *   HMAC方式によるJWT符号化・復号化のサンプル
 * * [JWT(+JWS)について調べた - takapiのブログ](https://takapi86.hatenablog.com/entry/2017/10/22/214703)
 *   JWT生成サンプル
 * * [How to test authenticate JWT routes in go - Stack Overflow](https://stackoverflow.com/questions/56415581/how-to-test-authenticate-jwt-routes-in-go)
 *   JWTの検証コードサンプル
 * * [JWT(JSON Web Token)を使った認証を試みる | 69log](https://blog.kazu69.net/2016/07/30/authenticate_with_json_web_token/)
 *   JSONを電子署名したurl-safe(URLで使用できない文字が含まれる)なclaimのことを指す
 * * [golang-jwt/jwt： Community maintained clone of https：//github.com/dgrijalva/jwt-go](https://github.com/golang-jwt/jwt/tree/main)
 *   上記以外はソースの引数の型と戻り値の型及びテストコードからルーチンの利用方法を読み取った。
 */

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
//jwtパッケージは公開鍵暗号のRSA鍵ファイルの読込みなどを想定しパーサに関数を渡す仕様
func MakeTokenSecretSender(secret[]byte) func(token *jwt.Token) (interface{}, error) {
  secretSender := func(token *jwt.Token) (interface{}, error) {
    return secret, nil
  }
  return secretSender
}



