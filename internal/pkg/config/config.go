package config //担当：石森

import (
  "os"
  "github.com/go-ozzo/ozzo-validation/v4"
  "gopkg.in/yaml.v2"
  "optim_22_app/pkg/log"
)

//Configは、アプリケーションの設定を表す
type Config struct {
  //サーバのドメイン
  Domain                 string `yaml:"domain"`
  //サーバーのポート
  ServerPort             int    `yaml:"server_port"`
  //データベースに接続するためのデータソース名(DSN)
  DSN                    string `yaml:"dsn"`
  //JWTの有効期限
  RefreshTokenExpiration int    `yaml:"refresh_token_expiration"`
  AccessTokenExpiration  int    `yaml:"access_token_expiration"`
  //JWTの署名キー
  RefreshTokenSecret     string `yaml:"refresh_token_secret"`
  AccessTokenSecret      string `yaml:"access_token_secret"`
  //HTTP通信における各種タイムアウト設定
  ReadTimeout            int64  `yaml:"read_timeout"`
  ReadHeaderTimeout      int64  `yaml:"read_header_timeout"`
  WriteTimeout           int64  `yaml:"write_timeout"`
  IdleTimeout            int64  `yaml:"idle_timeout"`
}

//アプリケーションの設定を検証
func (c Config) Validate() error {
  return validation.ValidateStruct(&c,
    validation.Field(&c.ServerPort, validation.Required),
    validation.Field(&c.DSN, validation.Required),
    validation.Field(&c.RefreshTokenExpiration, validation.Required),
    validation.Field(&c.AccessTokenExpiration, validation.Required),
    validation.Field(&c.RefreshTokenSecret, validation.Required),
    validation.Field(&c.ReadTimeout, validation.Required),
    validation.Field(&c.ReadHeaderTimeout, validation.Required),
    validation.Field(&c.WriteTimeout, validation.Required),
    validation.Field(&c.IdleTimeout, validation.Required),
  )
}

//設定ファイルから入力されたアプリケーション設定を返す
func Load(file string, logger log.Logger) (*Config, error) {
  
  var c Config

  //#region 設定ファイル`./configs/app.yaml`からConfigを設定、返す
  bytes, err := os.ReadFile(file)

  if err != nil {
    return nil, err
  }

  if err = yaml.Unmarshal(bytes, &c); err != nil {
    return nil, err
  }

  //検証
  if err = c.Validate(); err != nil {
    return nil, err
  }

  return &c, err
  //#endregion
}