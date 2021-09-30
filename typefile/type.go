// type.goに利用する構造体をまとめる。
// 構造体の属性は必ず最初の文字を大文字にする。
package typefile

import (
	"time"
)

type User struct{
	// 「"user_id": ユーザID」はすべてユーザがユーザ登録時に自身で設定したユーザIDのことであり，
	// サーバ側でデータベースに登録する際にインクリメンタルで付与される識別IDとは異なる。
	// なので、後で変更する。
	ID             int            `gorm:"primaryKey",json:"user_id"`
	Name           string         `gorm:"not null",json:"username"`
}

type Client struct{
	User           User           `gorm:"embedded"`
}

type Engineer struct{
	User           User           `gorm:"embedded"`
	Requests       []Request      `gorm:"many2many:engineer_requests;"`
}

type Winner struct{
	EngineerID     int            `gorm:"not null",json:"engineer_id"`
	Engineer       Engineer
	RequestID      int            `gorm:"unique;not null",json:"request_id"`
}

type Request struct{
	ID             int            `gorm:"primaryKey",json:"request_id"`
	ClientID       int            `gorm:"not null",json:"client_id"`
	Client         Client
	Engineers      []Engineer     `gorm:"many2many:engineer_requests;"`
	RequestName    string         `gorm:"not null",json:"requestname"`
	Content        string         `gorm:"not null",json:"content"`
	Winner         Winner
	Finish         bool           `gorm:"not null",json:"finish"`
	CreatedAt      time.Time      `gorm:"not null",json:"createdat"`
	UpdatedAt      time.Time      `gorm:"not null"`
}

type Submission struct{
	ID             int            `gorm:"primaryKey",json:"submission_id"`
	RequestID      int            `gorm:"not null",json:"request_id"`
	Request        Request
	EngineerID     int            `gorm:"not null",json:"engineer_id"`
	Engineer       Engineer
	Content        string         `gorm:"not null",json:"content"`
	CreatedAt      time.Time      `gorm:"not null",json:"createdat"`
	UpdatedAt      time.Time      `gorm:"not null"`
}