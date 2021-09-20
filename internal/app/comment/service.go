package comment

//#region コメント
type Comment struct {
  Id         int    `json:"id"`
  UserID     int    `json:"userID"`
  Date       time.Time `json:"date"`
  Title      string    `json:"title"`
  Text       string    `json:"text"`
  ReplyID    int    `json:"replyID"`
  Attachment []byte    `json:"attachment"`
}


func (m Comment) Validate() error {
  return validation.ValidateStruct(&m,
    validation.Field(&m.UserID, validation.Required, is.Int),
    validation.Field(&m.Date, validation.Required, validation.Date("2006-01-02")),
    validation.Field(&m.Title, validation.Required, validation.Length(3, 64)),
    validation.Field(&m.Text, validation.Required, validation.Length(3, 128)),
    validation.Field(&m.ReplyID, validation.Required, is.Int),
    //validation.Field(&m.Attachment, validation.Length(3, 128)),
  )
}
//#endregion




