package engineer

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"optim_22_app/model"
	"typefile"
	"strconv"
)

// JoinRequestで得たデータによって、エンジニアが特定リクエストに参加することをデータベースに登録する。
func CreateEngineerJoin(c *gin.Context) {
	// formから送られた値を得る
    request_id_string := c.PostForm("request_id")
    engineer_id_string := c.PostForm("engineer_id")

	// 文字列をintに変換
	request_id, _ := strconv.Atoi(request_id_string)
	engineer_id, _ := strconv.Atoi(engineer_id_string)

	// Request構造体を格納するための型を宣言
	request := typefile.Request{}
	// Engineer構造体を格納するための型を宣言
	engineer := typefile.Engineer{}

	model.Db.Find(&engineer,"id = ?",engineer_id)
	model.Db.Find(&request,"id = ?",request_id)
	// many2manyのtableにデータを格納する
	model.Db.Model(&engineer).Association("Requests").Append(&request)

	// StatusSeeOther = 303,違うコンテンツだけどリダイレクト
    c.Redirect(http.StatusSeeOther, "//localhost:8080/")
}

// エンジニアがsubmissionを提出するためのページを表示する
func NewSubmission(c *gin.Context) {
	// urlの引数で受け取ったrequest_idをrequest_id_stringという変数に格納している。
	request_id_string := c.Param("request_id")
	// 文字列をintに変換
	request_id, _ := strconv.Atoi(request_id_string)

	engineer_id := 3

    c.HTML(http.StatusOK, "new_submission.html", gin.H{
    	"engineer_id": engineer_id,
    	"request_id": request_id,
    })
}

// NewSubmissionで得たデータによって、エンジニアが提出した提出物をデータベースに登録する。
func CreateSubmission(c *gin.Context) {
	// formから送られた値を得る
    request_id_string := c.PostForm("request_id")
    engineer_id_string := c.PostForm("engineer_id")
    content := c.PostForm("Content")

	// 文字列をintに変換
	request_id, _ := strconv.Atoi(request_id_string)
	engineer_id, _ := strconv.Atoi(engineer_id_string)

	// EngineerIDはUser機能が作成された後に、IDの取得方法を聞いた後に変更する。
    var submission = typefile.Submission{RequestID: request_id,EngineerID: engineer_id,Content: content}
    model.Db.Create(&submission)

	// StatusSeeOther = 303,違うコンテンツだけどリダイレクト
    c.Redirect(http.StatusSeeOther, "//localhost:8080/")
}

// エンジニアが参加しているリクエストを表示する
func ShowJoinRequest(c *gin.Context) {
	// urlの引数で受け取ったengineer_idをengineer_id_stringという変数に格納している。
	engineer_id_string := c.Param("engineer_id")
	// 文字列をintに変換
	engineer_id, _ := strconv.Atoi(engineer_id_string)

	// Engineer構造体データを格納するためのインスタンスを生成
	engineer := typefile.Engineer{}
	// Request構造体データを複数格納するためのインスタンスを生成
	requests := []typefile.Request{}

	model.Db.Find(&engineer,"id = ?",engineer_id)
	model.Db.Model(&engineer).Association("Requests").Find(&requests)

    c.HTML(http.StatusOK, "show_join_request.html", gin.H{
    	"requests": requests,
    	"engineer": engineer,
    })
}

// エンジニアが提出済みのsubmissionを編集する。
func EditSubmission(c *gin.Context) {
	// urlの引数で受け取ったrequest_idをrequest_idという変数に格納している。
	submission_id_string := c.Param("submission_id")
	// 文字列をintに変換
	submission_id, _ := strconv.Atoi(submission_id_string)

	// Submission構造体データを格納するためのインスタンスを生成
	submission := typefile.Submission{}

	model.Db.Find(&submission,"id = ?",submission_id)

	c.HTML(http.StatusOK, "edit_submission.html", gin.H{
		"submission": submission,
	})
}

// エンジニアが編集済みのsubmissionを提出する。
func UpdateSubmission(c *gin.Context) {
	// formから送られた値を得る
	submission_id_string := c.PostForm("submission_id")
    content := c.PostForm("Content")

	// 文字列をintに変換
	submission_id, _ := strconv.Atoi(submission_id_string)

	// Submission構造体データを格納するためのインスタンスを生成
	submission := typefile.Submission{}

	// 該当するsubmissionを抽出している。
	model.Db.Find(&submission,"id = ?",submission_id)

	// contentを更新する。
	submission.Content = content
	model.Db.Save(&submission)

	// StatusSeeOther = 303,違うコンテンツだけどリダイレクト
    c.Redirect(http.StatusSeeOther, "//localhost:8080/")
}