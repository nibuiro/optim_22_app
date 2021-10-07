package comment


func Presenter(output []byte) string {
  return `{"comments":`+string(output[:])+"}"
}