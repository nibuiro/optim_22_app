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


//#endregion




