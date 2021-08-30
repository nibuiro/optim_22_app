package config

import (
  "os"
  "github.com/go-ozzo/ozzo-validation/v4"
  "gopkg.in/yaml.v2"
  "optim_22_app/pkg/log"
)

//Configは、アプリケーションの設定を表す
type Config struct {
  //サーバーのポート
  ServerPort int `yaml:"server_port"`
  // データベースに接続するためのデータソース名(DSN)
  DSN string `yaml:"dsn"`
  //JWTの署名キー
  JWTSigningKey string `yaml:"jwt_signing_key"`
}

//アプリケーションの設定を検証
func (c Config) Validate() error {
  return validation.ValidateStruct(&c,
    validation.Field(&c.ServerPort, validation.Required),
    validation.Field(&c.DSN, validation.Required),
    validation.Field(&c.JWTSigningKey, validation.Required),
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