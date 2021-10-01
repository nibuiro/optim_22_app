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
    request_id_string := c.Query("request_id")
    engineer_id_string := c.Query("engineer_id")

	// 文字列をintに変換
	request_id, _ := strconv.Atoi(request_id_string)
	engineer_id, _ := strconv.Atoi(engineer_id_string)

	// Request構造体を格納するための型を宣言
	request := typefile.Request{}
	// Engineer構造体を格納するための型を宣言
	engineer := typefile.Engineer{}

	model.Db.Find(&engineer,"id = ?",engineer_id)
	model.Db.Find(&request,"id = ?",request_id)
	// エンジニアが参加しているリクエストを外部キーなしで取得するために、Associationを追加している。
	model.Db.Model(&engineer).Association("Requests").Append(&request)
	// リクエストに参加しているエンジニアを外部キーなしで取得するために、Associationを追加している。
	model.Db.Model(&request).Association("Engineers").Append(&engineer)

	// StatusSeeOther = 303,違うコンテンツだけどリダイレクト
    c.Redirect(http.StatusSeeOther, "//localhost:8080/")
}

// NewSubmissionで得たデータによって、エンジニアが提出した提出物をデータベースに登録する。
func CreateSubmission(c *gin.Context) {
	// urlのクエリパラメータで受け取ったrequest_idをrequest_id_stringという変数に格納している。
	request_id_string := c.Query("request_id")
	// 文字列をintに変換
	request_id, _ := strconv.Atoi(request_id_string)
	// urlのクエリパラメータで受け取ったengineer_idをengineer_id_stringという変数に格納している。
	engineer_id_string := c.Query("engineer_id")
	// 文字列をintに変換
	engineer_id, _ := strconv.Atoi(engineer_id_string)

	// フロントから送られたrequestのjsonデータをバインドするための構造体を宣言
	var submission typefile.Submission

	// JSONからrequest構造体へ値をマッピングしている。
	if err := c.ShouldBindJSON(&submission); err != nil {
		// エラーが生じた場合、内容を出力。ない場合は、何も出力しない。
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// データを追加している。
	submission.EngineerID = engineer_id
	submission.RequestID = request_id

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
	// Associationによって、requestデータを取り出す。
	model.Db.Model(&engineer).Association("Requests").Find(&requests)

    c.HTML(http.StatusOK, "show_join_request.html", gin.H{
    	"requests": requests,
    	"engineer": engineer,
    })
}

// エンジニアが編集済みのsubmissionを提出する。
func UpdateSubmission(c *gin.Context) {
	// urlの引数で受け取ったsubmission_idをsubmission_id_stringという変数に格納している。
	submission_id_string := c.Query("submission_id")
	// 文字列をintに変換
	submission_id, _ := strconv.Atoi(submission_id_string)

	// フロントから送られたrequestのjsonデータをバインドするための構造体を宣言
	var submissionjson typefile.Submission
	// Submission構造体データを格納するためのインスタンスを生成
	submission := typefile.Submission{}

	// JSONからrequest構造体へ値をマッピングしている。
	if err := c.ShouldBindJSON(&submissionjson); err != nil {
		// エラーが生じた場合、内容を出力。ない場合は、何も出力しない。
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 該当するsubmissionを抽出している。
	model.Db.Find(&submission,"id = ?",submission_id)

	// contentを更新する。
	submission.Content = submissionjson.Content
	model.Db.Save(&submission)

	// StatusSeeOther = 303,違うコンテンツだけどリダイレクト
    c.Redirect(http.StatusSeeOther, "//localhost:8080/submission/show_submission/:submission_id")
}