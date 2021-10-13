package engineer

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"optim_22_app/model"
	"typefile"
	"strconv"
)

type ReceiveSubmissionJson struct{
	ID             int
	RequestID      string         `json:"request_id"`
	EngineerID     string         `json:"engineer_id"`
	URL            string         `json:"url"`
	Content        string         `json:"content"`
}

// フロントからサーバーサイドにrequest参加に関するjsonデータが送られた際に利用する構造体
type ReceiveRequestJoinJson struct{
	EngineerID     string         `json:"engineer_id"`
	RequestID      string         `json:"request_id"`
}

// JoinRequestで得たデータによって、エンジニアが特定リクエストに参加することをデータベースに登録する。
func CreateEngineerJoin(c *gin.Context) {
	// フロントから送られたrequest参加に関するjsonデータをバインドするための構造体を宣言
	var request_join_json ReceiveRequestJoinJson

	// JSONからwinner構造体へ値をマッピングしている。
	if err := c.ShouldBindJSON(&request_join_json); err != nil {
		// エラーが生じた場合、内容を出力。ない場合は、何も出力しない。
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 文字列をintに変換
	request_id, _:= strconv.Atoi(request_join_json.RequestID)
	engineer_id, _ := strconv.Atoi(request_join_json.EngineerID)

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
    // c.Redirect(http.StatusSeeOther, "//localhost:8080/api/requests")
}

// NewSubmissionで得たデータによって、エンジニアが提出した提出物をデータベースに登録する。
func CreateSubmission(c *gin.Context) {
	// フロントから送られたrequestのjsonデータをバインドするための構造体を宣言
	var submission_json ReceiveSubmissionJson

	// JSONからrequest構造体へ値をマッピングしている。
	if err := c.ShouldBindJSON(&submission_json); err != nil {
		// エラーが生じた場合、内容を出力。ない場合は、何も出力しない。
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Submission構造体データを格納するためのインスタンスを生成
	submission := typefile.Submission{}

	// submission_jsonが持つデータをsubmissionのそれぞれの対応する属性に格納する。
	submission.EngineerID, _ = strconv.Atoi(submission_json.EngineerID)
	submission.RequestID, _ = strconv.Atoi(submission_json.RequestID)
	submission.URL = submission_json.URL
	submission.Content = submission_json.Content

    model.Db.Create(&submission)

	// StatusSeeOther = 303,違うコンテンツだけどリダイレクト
    c.Redirect(http.StatusSeeOther, "//localhost:8080/api/requests")
}

// エンジニアが編集済みのsubmissionを提出する。
func UpdateSubmission(c *gin.Context) {
	// urlの引数で受け取ったsubmission_idをsubmission_id_stringという変数に格納している。
	submission_id_string := c.Param("submission_id")
	// 文字列をintに変換
	submission_id, _ := strconv.Atoi(submission_id_string)

	// フロントから送られたrequestのjsonデータをバインドするための構造体を宣言
	var submission_json ReceiveSubmissionJson
	// Submission構造体データを格納するためのインスタンスを生成
	submission := typefile.Submission{}

	// JSONからrequest構造体へ値をマッピングしている。
	if err := c.ShouldBindJSON(&submission_json); err != nil {
		// エラーが生じた場合、内容を出力。ない場合は、何も出力しない。
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 該当するsubmissionを抽出している。
	model.Db.Find(&submission,"id = ?",submission_id)

	// contentを更新する。
	submission.Content = submission_json.Content
	submission.URL = submission_json.URL
	model.Db.Save(&submission)

	// redirect先を追加している。
	redirect_url := "//localhost:8080/api/submission/" + submission_id_string
    // StatusSeeOther = 303,違うコンテンツだけどリダイレクト
    c.Redirect(http.StatusSeeOther, redirect_url)
}