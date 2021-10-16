package roundary //repository boundary  so roundary :)

//モデル定義群のフォーク

import (
	"time"
)

type User struct{
	ID             int            `json:"user_id"`
	Name           string         `json:"username"`
	Email          string         `json:"email"`  
//	Password       string         `json:"password"`   セキュリティのため
}

type Profile struct{
	ID             int            `gorm:"primaryKey";autoIncrement:false`
	Bio            string        
	Sns            []byte         `json:"SNS"`
	Icon           string
}

type Client struct{
	ID             int            `json:"user_id"`
	Name           string         `json:"username"`
	Email          string         `json:"email"`  
	Bio            string         `json:"comment"`     
	Sns            string         `json:"SNS"`
	Icon           string         `json:"icon"`
}

type Engineer struct{
	ID             int            `json:"user_id"`
	Name           string         `json:"username"`
	Email          string         `json:"email"`  
	Bio            string         `json:"comment"`     
	Sns            string         `json:"SNS"`
	Icon           string         `json:"icon"`
}

type Winner struct{
	EngineerID     int            `json:"engineer_id"`
	RequestID      int            `json:"request_id"`
}

type Request struct{
	ID             int            `json:"request_id"`
	ClientID       int            `json:"client_id"`
	RequestName    string         `json:"requestname"`
	Content        string         `json:"content"`
	Finish         bool           `json:"finish"`
	CreatedAt      time.Time      `json:"createdat"`
	UpdatedAt      time.Time      `json:"updatedat"`
//	Client         Client         `json:"client"`      自身のため
	Winner         Winner         `json:"winner"`
	Engineers      []Engineer     `json:"engineers" gorm:"many2many:engineer_requests;"`
	Submission     []Submission   `json:"submission"`
}

type Submission struct{
	ID             int            `json:"submission_id"`
	RequestID      int            `json:"request_id"`
	EngineerID     int            `json:"engineer_id"`
	URL            string         `json:"url"`
	Content        string         `json:"content"`
	CreatedAt      time.Time      `json:"createdat"`
	UpdatedAt      time.Time      `json:"updatedat"`
//	Request        Request        `json:"request"      仕様
//	Engineer       Engineer       `json:"engineer"`    仕様
}

type Comment struct {
  CreatedAt  time.Time
  Id         int                 `gorm:"primaryKey"`
  RequestID  int                 `gorm:"not null"`
  UserID     int                 `gorm:"not null"`
  Title      string    
  Body       string    
  ReplyID    int  
}