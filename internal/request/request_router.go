package request

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"optim_22_app/model"
	"typefile"
	"strconv"
	"time"
)

//データを送信する際に利用する構造体を定義
type SubmissionJson struct{
	ID             int                 `json:"submission_id"`
	CreatedAt      time.Time           `json:"createdat"`
	// 要件はエンジニアのプロフィールデータであるが、プロフィール機能は担当外のため、EngineerIDを代用する。
	EngineerID     int                 `json:"engineer`
	Content        string              `json:"content"`
}

type RequestJson struct{
	ID             int                 `json:"request_id"`
	RequestName    string              `json:"requestname"`
	CreatedAt      time.Time           `json:"createdat"`
	// 要件はクライアントのプロフィールデータであるが、プロフィール機能は担当外のため、ClientIDを代用する。
	ClientID       int                 `json:"client"`
	// 要件はエンジニアのプロフィールデータであるが、プロフィール機能は担当外のため、EngineerIDを代用する。
	EngineersID    []int               `json:"engineers"`
	Content        string              `json:"content"`
	Submissions    []SubmissionJson    `json:"submissions"`
}

// 特定リクエストの詳細を表示する
func ShowRequest(c *gin.Context) {
	// urlの引数で受け取ったengineer_idをengineer_id_stringという変数に格納している。
	request_id_string := c.Param("request_id")
	// 文字列をintに変換
	request_id, _ := strconv.Atoi(request_id_string)

	// Request構造体を格納するためのインスタンスを生成
	request := typefile.Request{}
	// Winner構造体を格納するためのインスタンスを生成
	winner := typefile.Winner{}
	// Engineer構造体を複数格納するためのインスタンスを生成
	engineers := []typefile.Engineer{}
	// Submission構造体を複数格納するためのインスタンスを生成
	submissions := []typefile.Submission{}
	// RequestJson構造体データを格納するためのインスタンスを生成
	request_json := RequestJson{}
	
	// 特定のidを持つRequestを抽出する。
	model.Db.Find(&request,"id = ?",request_id)
	// SELECT * FROM `requests` WHERE id = ?
	// 特定のrequest_idを持つwinnerを抽出する。
	model.Db.Find(&winner,"request_id = ?",request_id)
	// SELECT * FROM `winners` WHERE request_id = ?
	// Associationによって、engineerデータを取り出す。
	model.Db.Model(&request).Association("Engineers").Find(&engineers)
	// 特定のrequest_idを持つsubmissionを全抽出する。
	model.Db.Find(&submissions,"request_id = ?",request.ID)
	// SELECT * FROM `submissions` WHERE request_id = ?

	// requestが持つデータをrequest_jsonのそれぞれの対応する属性に格納する。
	request_json.ID = request.ID
	request_json.RequestName = request.RequestName
	request_json.CreatedAt = request.CreatedAt
	request_json.ClientID = request.ClientID
	request_json.Content = request.Content

	// 抽出したengineersデータからループ処理でエンジニアidを取得し、engineers_id配列に格納している。
	for _, engineer := range engineers{
		request_json.EngineersID = append(request_json.EngineersID,engineer.User.ID)
	}

	// submissionは複数存在するため、submissionデータをループで追加していく。
	for _, submission := range submissions{
		// SubmissionJson構造体データを格納するためのインスタンスを生成
		submission_json := SubmissionJson{}

		// submissionが持つデータをsubmission_jsonのそれぞれの対応する属性に格納する。
		submission_json.ID = submission.ID
		submission_json.CreatedAt = submission.CreatedAt
		submission_json.EngineerID = submission.EngineerID
		submission_json.Content = submission.Content
		request_json.Submissions = append(request_json.Submissions,submission_json)
	}

    c.JSON(http.StatusOK, gin.H{
    	"request": request_json,
    })
}