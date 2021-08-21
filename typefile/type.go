// type.goに利用する構造体をまとめる。
// 構造体の属性は必ず最初の文字を大文字にする。

package typefile

// EngineerID  []int 保留
type Request struct{
	ID          uint           `gorm:"primaryKey"`
	ClientId    int
	Name        string
	Content     string
	WinnerId    int
	Finish      bool
}