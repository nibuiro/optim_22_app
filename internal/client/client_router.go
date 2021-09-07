package client

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"optim_22_app/model"
	"typefile"
	"strconv"
)

// クライアントがリクエストを依頼する。(入力画面の表示)
func NewRequest(c *gin.Context) {
    c.HTML(http.StatusOK, "new_request.html", gin.H{})
}

// クライアントがリクエストを依頼する。(入力内容をDBに格納)
func CreateRequest(c *gin.Context) {
	// formで入力された値を得る
    requestname := c.PostForm("RequestName")
    content := c.PostForm("Content")
    // ClientIDはUser機能が作成された後に、IDの取得方法を聞いた後に変更する。
    var request = typefile.Request{ClientID: 6,RequestName: requestname,Content: content,Finish: false}
    model.Db.Create(&request)

    // StatusSeeOther = 303,違うコンテンツだけどリダイレクト
    c.Redirect(http.StatusSeeOther, "//localhost:8080/hello")
}

// クライアントが依頼したリクエストの一覧を表示する。
// クライアントがリクエストを依頼する。(入力画面の表示)
func ShowRequest(c *gin.Context) {
	// urlの引数で受け取ったclient_idをclient_idという変数に格納している。
	client_id_string := c.Param("client_id")
	// 文字列をintに変換
	client_id, _ := strconv.Atoi(client_id_string)
	// Request構造体を複数格納するための型を宣言
	requests := []typefile.Request{}
	// 特定のclient_idを持つリクエストを全抽出する。ClientIDはUser機能が作成された後に、IDの取得方法を聞いた後に変更する。
	model.Db.Find(&requests,"client_id = ?",client_id)
	// SELECT * FROM `requests` WHERE client_id = ?

    c.HTML(http.StatusOK, "show_request.html", gin.H{
    	"requests": requests,
    	"client_id": client_id,
    })
}

// 特定リクエストのサブミッションの一覧を標示するための関数
func ShowSubmission(c *gin.Context) {
	// Submission構造体を複数格納するための型を宣言
	submissions := []typefile.Submission{}
	// urlの引数で受け取ったrequest_idをrequest_idという変数に格納している。
	request_id_string := c.Param("request_id")
	// 文字列をintに変換
	request_id, _ := strconv.Atoi(request_id_string)
	// 特定のrequest_idを持つsubmissionを全抽出する。
	model.Db.Find(&submissions,"request_id = ?",request_id)
	// SELECT * FROM `requests` WHERE request_id = ?

    c.HTML(http.StatusOK, "show_submission.html", gin.H{
    	"submissions": submissions,
    	"request_id": request_id,
    })

    // winnerがいない場合はwinnerがいないと表示し、winnerがいる場合は～さんがwinnerですと表示する。
}

// 勝者を決定するための関数
func DecideWinner(c *gin.Context) {
	// formから送られた値を得る
    request_id_string := c.PostForm("request_id")
    engineer_id_string := c.PostForm("engineer_id")

	// 文字列をintに変換
	request_id, _ := strconv.Atoi(request_id_string)
	engineer_id, _ := strconv.Atoi(engineer_id_string)
    // ClientIDはUser機能が作成された後に、IDの取得方法を聞いた後に変更する。
    var winner = typefile.Winner{EngineerID: engineer_id,RequestID: request_id}
    model.Db.Create(&winner)

    // StatusSeeOther = 303,違うコンテンツだけどリダイレクト
    c.Redirect(http.StatusSeeOther, "//localhost:8080/hello")
}