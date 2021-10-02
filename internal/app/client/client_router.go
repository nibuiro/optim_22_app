package client

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"optim_22_app/model"
	"typefile"
	"strconv"
)

// フロントからサーバーサイドにrequestのjsonデータが送られた際に利用する構造体
type ReceiveRequestJson struct{
	ID             int
	RequestName    string              `json:"requestname"`
	// 要件はクライアントのプロフィールデータであるが、プロフィール機能は担当外のため、ClientIDを代用する。
	ClientID       string              `json:"client_id"`
	// 要件はエンジニアのプロフィールデータであるが、プロフィール機能は担当外のため、EngineerIDを代用する。
	Content        string              `json:"content"`
	Finish         bool                `gorm:"not null",json:"finish"`
}

// フロントからサーバーサイドにwinnerのjsonデータが送られた際に利用する構造体
type ReceiveWinnerJson struct{
	EngineerID     string              `gorm:"not null",json:"engineer_id"`
	RequestID      string              `gorm:"unique;not null",json:"request_id"`
}

// クライアントがリクエストを依頼する。(入力内容をDBに格納)
func CreateRequest(c *gin.Context) {
	// フロントから送られたrequestのjsonデータをバインドするための構造体を宣言
	var request_json ReceiveRequestJson
	// Request構造体を格納するためのインスタンスを生成
	request := typefile.Request{}

	// JSONからrequest構造体へ値をマッピングしている。
	if err := c.ShouldBindJSON(&request_json); err != nil {
		// エラーが生じた場合、内容を出力。ない場合は、何も出力しない。
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// request_jsonが持つデータをrequestのそれぞれの対応する属性に格納する。
	request.RequestName = request_json.RequestName
	request.ClientID, _ = strconv.Atoi(request_json.ClientID)
	request.Content = request_json.Content
	request.Finish = request_json.Finish

	c.JSON(http.StatusOK, gin.H{
    	"request": request,
    })
    
    //model.Db.Create(&request)

    // 依頼作成者とログインユーザーが一致していることを確かめている。
    // if request.ClientID == user_id{
    	// model.Db.Create(&request)
        // StatusSeeOther = 303,違うコンテンツだけどリダイレクト
        // c.Redirect(http.StatusSeeOther, "//localhost:8080/")
    // }else{
    	// StatusSeeOther = 303,違うコンテンツだけどリダイレクト
        // c.Redirect(http.StatusSeeOther, "//localhost:8080/")
    // }
    // StatusSeeOther = 303,違うコンテンツだけどリダイレクト
    //c.Redirect(http.StatusSeeOther, "//localhost:8080/")
}

// クライアントが依頼したリクエストの一覧を表示する。(すべてのユーザーが特定クライアントのリクエスト一覧を表示できる。)
func ShowRequest(c *gin.Context) {
	// urlのクエリパラメータで受け取ったclient_idをclient_id_stringという変数に格納している。
	client_id_string := c.Param("client_id")
	// 文字列をintに変換
	client_id, _ := strconv.Atoi(client_id_string)

	// Request構造体を複数格納するためのインスタンスを生成
	requests := []typefile.Request{}
	// Client構造体を格納するためのインスタンスを生成
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
func UpdateRequest(c *gin.Context) {
	// urlの引数で受け取ったrequest_idをrequest_id_stringという変数に格納している。
	request_id_string := c.Param("request_id")
	// 文字列をintに変換
	request_id, _ := strconv.Atoi(request_id_string)
	// フロントから送られたrequestのjsonデータをバインドするための構造体を宣言
	var request_json ReceiveRequestJson
	// Request構造体を格納するためのインスタンスを生成
	request := typefile.Request{}

	// JSONからrequest構造体へ値をマッピングしている。
	if err := c.ShouldBindJSON(&request_json); err != nil {
		// エラーが生じた場合、内容を出力。ない場合は、何も出力しない。
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 該当するリクエストを抽出している。
	model.Db.Find(&request,"id = ?",request_id)

	// それぞれのcolumeの値を更新する。
	request.RequestName = request_json.RequestName
	request.Content = request_json.Content
	model.Db.Save(&request)

    // リクエストが終了していないことを確かめる。
    // if request.Finish == false{
    	// StatusSeeOther = 303,違うコンテンツだけどリダイレクト
    	//c.Redirect(http.StatusSeeOther, "//localhost:8080/show_request/:request_id")
    //}

    // 依頼者と依頼更新者が一致していることを確かめている。
    // if update_request.ClientID == user_id{
    	// model.Db.Find(&request,"id = ?",request_id).Updates(&update_request)
        // StatusSeeOther = 303,違うコンテンツだけどリダイレクト
        // c.Redirect(http.StatusSeeOther, "//localhost:8080/")
    // }else{
    	// StatusSeeOther = 303,違うコンテンツだけどリダイレクト
        // c.Redirect(http.StatusSeeOther, "//localhost:8080/")
    // }

    // redirect先を追加している。
    redirect_url := "//localhost:8080/request/show_request/" + request_id_string
    // StatusSeeOther = 303,違うコンテンツだけどリダイレクト
    c.Redirect(http.StatusSeeOther, redirect_url)
}

// 特定リクエストのサブミッションの一覧を標示するための関数
func ShowSubmission(c *gin.Context) {
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
}

// 勝者を決定するための関数
func DecideWinner(c *gin.Context) {
	// フロントから送られたwinnerのjsonデータをバインドするための構造体を宣言
	var winner_json ReceiveWinnerJson

	// JSONからwinner構造体へ値をマッピングしている。
	if err := c.ShouldBindJSON(&winner_json); err != nil {
		// エラーが生じた場合、内容を出力。ない場合は、何も出力しない。
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 文字列をintに変換
	request_id, _ := strconv.Atoi(winner_json.RequestID)
	engineer_id, _ := strconv.Atoi(winner_json.EngineerID)

    // ClientIDはUser機能が作成された後に、IDの取得方法を聞いた後に変更する。
    var winner = typefile.Winner{EngineerID: engineer_id,RequestID: request_id}
    model.Db.Create(&winner)

    // Request構造体を格納するためのインスタンスを生成
	request := typefile.Request{}

	// 該当するリクエストを抽出している。
	model.Db.Find(&request,"id = ?",request_id)

	// Finishをtrueに更新する。
	request.Finish = true
	model.Db.Save(&request)

    // 依頼者とwinner決定者が一致していることを確かめている
    // if request.ClientID == user_id{
    	// var winner = typefile.Winner{EngineerID: engineer_id,RequestID: request_id}
        // model.Db.Create(&winner)
        // StatusSeeOther = 303,違うコンテンツだけどリダイレクト
        // c.Redirect(http.StatusSeeOther, "//localhost:8080/")
    // }else{
    	// StatusSeeOther = 303,違うコンテンツだけどリダイレクト
        // c.Redirect(http.StatusSeeOther, "//localhost:8080/")
    // }

    // StatusSeeOther = 303,違うコンテンツだけどリダイレクト
    c.Redirect(http.StatusSeeOther, "//localhost:8080/")
}