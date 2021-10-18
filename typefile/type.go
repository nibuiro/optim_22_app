// type.goに利用する構造体をまとめる。
// 構造体の属性は必ず最初の文字を大文字にする。
package typefile

import (
	"time"
)

type User struct{
	ID             int            `gorm:"primaryKey",json:"user_id"`
	Name           string         `gorm:"not null",json:"username"`
	Email          string         `gorm:"type:varchar(100);unique",json:"email"`  
	Password       string         `gorm:"not null",json:"password"`     
}

type Profile struct{
	ID             int            `gorm:"primaryKey"`//;autoIncrement:false`
	Bio            string         
	Sns            []byte         
	Icon           string         
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
	URL            string         `gorm:"not null",json:"url"`
	Content        string         `gorm:"not null",json:"content"`
	CreatedAt      time.Time      `gorm:"not null",json:"createdat"`
	UpdatedAt      time.Time      `gorm:"not null"`
}

type Comment struct {
  CreatedAt  time.Time
  Id         int                 `gorm:"primaryKey"`
  RequestID  int                 `gorm:"not null"`
  UserID     int                 `gorm:"not null"`
  Title      string    
  Body       string    
  ReplyID    int  
  Attachment string              `gorm:"type:varchar(2048)"`  
}