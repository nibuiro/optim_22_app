package home

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"optim_22_app/model"
	"typefile"
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

type RequestsJson struct{
	Requests       []RequestJson       `json:"requests"`
}

// ホームページを表示する
func ShowHomepage(c *gin.Context) {
	// Request構造体データを複数格納するためのインスタンスを生成
	requests := []typefile.Request{}
	// RequestJson構造体データを格納するためのインスタンスを生成
	requests_json := RequestsJson{}

	//　受付中のリクエストを格納する。(Winnerを持たないリクエストを抽出する)
	model.Db.Find(&requests,"finish = ?",0)
	// SELECT * FROM `requests` WHERE finish = ?

	for _, request := range requests{
		// RequestJson構造体データを格納するためのインスタンスを生成
		request_json := RequestJson{}
		// Submission構造体データを複数格納するためのインスタンスを生成
		submissions := []typefile.Submission{}
		// Engineer構造体を複数格納するためのインスタンスを生成
		engineers := []typefile.Engineer{}

		//　特定リクエストのidを持つsubmissionを格納する。
		model.Db.Find(&submissions,"request_id = ?",request.ID)
		// SELECT * FROM `submissions` WHERE request_id = ?
		// Associationによって、engineerデータを取り出す。
		model.Db.Model(&request).Association("Engineers").Find(&engineers)

		// requestが持つデータをrequest_jsonのそれぞれの対応する属性に格納する。
		request_json.ID = request.ID
		request_json.RequestName = request.RequestName
		request_json.CreatedAt = request.CreatedAt
		request_json.ClientID = request.ClientID
		request_json.Content = request.Content
		// EngineersIDは複数存在するため、エンジニアのIDをループで追加していく。
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
		requests_json.Requests = append(requests_json.Requests,request_json)
	}

	c.JSON(http.StatusOK, gin.H{
    	"requests": requests_json.Requests,
    })
}