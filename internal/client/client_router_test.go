package client

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

// assert.Equal(t, expected, actual)

// クライアントがリクエストを依頼する。(入力内容をDBに格納)
func TestCreateRequest(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	// ログインユーザーのidをクライアントidとし、リクエスト作成が成功するか確かめる。
	t.Run("success CreateRequest()", func(t *testing.T){
		var requestname string = "requestname1"
		var content string = "content1"
		var client_id int = 1
		// ClientIDはUser機能が作成された後に、IDの取得方法を聞いた後に変更する。
		var request = typefile.Request{ClientID: client_id,RequestName: requestname,Content: content}
		if err := model.Db.Create(&request).Error; err != nil {
			t.Errorf("failed CreateRequest()")
		}

		assert.Equal(t, client_id, request.ClientID)
		assert.Equal(t, requestname, request.RequestName)
		assert.Equal(t, email, user.Email)

		t.Logf("request: %p", request)
		t.Logf("request.ClientID: %d", request.ClientID)
		t.Logf("request.RequestName: %s", request.RequestName)
		t.Logf("request.Content: %s", request.Content)
	}

	// ログインユーザーではないidをクライアントidとし、リクエスト作成が失敗するか確かめる。
	t.Run("unsuccess CreateRequest()", func(t *testing.T){
		var requestname string = "requestname2"
		var content string = "content2"
		var client_id int = 2
		// ClientIDはUser機能が作成された後に、IDの取得方法を聞いた後に変更する。
		var request = typefile.Request{ClientID: client_id,RequestName: requestname,Content: content}
		if err := model.Db.Create(&request).Error; err != nil {
			t.Errorf("failed CreateRequest()")
		}

		assert.Equal(t, client_id, request.ClientID)
		assert.Equal(t, requestname, request.RequestName)
		assert.Equal(t, email, user.Email)

		t.Logf("request: %p", request)
		t.Logf("request.ClientID: %d", request.ClientID)
		t.Logf("request.RequestName: %s", request.RequestName)
		t.Logf("request.Content: %s", request.Content)
	}
}

// クライアントが依頼したリクエストの一覧を表示する。
// クライアントがリクエストを依頼する。(入力画面の表示)
func TestShowRequest(t *testing.T) {
	// urlの引数で受け取ったclient_idをclient_idという変数に格納している。
	client_id_string := c.Param("client_id")
	// 文字列をintに変換
	client_id, _ := strconv.Atoi(client_id_string)

	// Request構造体を複数格納するための型を宣言
	requests := []typefile.Request{}
	// Client構造体を格納するための型を宣言
	client := typefile.Client{}
	
	// 特定のclient_idを持つリクエストを全抽出する。ClientIDはUser機能が作成された後に、IDの取得方法を聞いた後に変更する。
	model.Db.Find(&requests,"client_id = ?",client_id)
	// SELECT * FROM `requests` WHERE client_id = ?
	// 特定のidを持つclientを抽出する。
	model.Db.Find(&client,"id = ?",client_id)
	// SELECT * FROM `clients` WHERE id = ?

    c.HTML(http.StatusOK, "show_request.html", gin.H{
    	"requests": requests,
    	"client": client,
    })
}

// クライアントが依頼済みのリクエストを編集する。
func TestEditRequest(t *testing.T) {
	// urlの引数で受け取ったrequest_idをrequest_idという変数に格納している。
	request_id_string := c.Param("request_id")
	// 文字列をintに変換
	request_id, _ := strconv.Atoi(request_id_string)

	c.HTML(http.StatusOK, "edit_request.html", gin.H{
		"request_id": request_id,
	})
}

// クライアントが依頼済みのリクエストを編集する。
func TestUpdateRequest(t *testing.T) {
	// formで入力された値を得る
    requestname := c.PostForm("RequestName")
    content := c.PostForm("Content")
    request_id_string := c.PostForm("request_id")

	// 文字列をintに変換
	request_id, _ := strconv.Atoi(request_id_string)

	// Request構造体を格納するための型を宣言
	request := typefile.Request{}
    // ClientIDはUser機能が作成された後に、IDの取得方法を聞いた後に変更する。
    var update_request = typefile.Request{ClientID: 3,RequestName: requestname,Content: content}

    model.Db.Find(&request,"id = ?",request_id).Updates(&update_request)

    // StatusSeeOther = 303,違うコンテンツだけどリダイレクト
    c.Redirect(http.StatusSeeOther, "//localhost:8080/hello")
}

// 特定リクエストのサブミッションの一覧を標示するための関数
func TestShowSubmission(t *testing.T) {
	// Submission構造体を複数格納するための型を宣言
	submissions := []typefile.Submission{}
	// Request構造体を格納するための型を宣言
	request := typefile.Request{}

	// urlの引数で受け取ったrequest_idをrequest_id_stringという変数に格納している。
	request_id_string := c.Param("request_id")
	// 文字列をintに変換
	request_id, _ := strconv.Atoi(request_id_string)

	// 特定のrequest_idを持つsubmissionを全抽出する。
	model.Db.Find(&submissions,"request_id = ?",request_id)
	// SELECT * FROM `submissions` WHERE request_id = ?
	// 特定のidを持つリクエストを抽出する。
	model.Db.Find(&request,"id = ?",request_id)
	// SELECT * FROM `requests` WHERE id = ?

    c.HTML(http.StatusOK, "show_submission.html", gin.H{
    	"submissions": submissions,
    	"request_id": request_id,
    })

    // winnerがいない場合はwinnerがいないと表示し、winnerがいる場合は～さんがwinnerですと表示する。
}

// 勝者を決定するための関数
func TestDecideWinner(t *testing.T) {
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