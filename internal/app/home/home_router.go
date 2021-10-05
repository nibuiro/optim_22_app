package home

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"optim_22_app/model"
	"typefile"
	"time"
)

//データを送信する際に利用する構造体を定義
type UserProfileJson struct{
	UserID            int                 `json:"user_id"`
	UserName          string              `json:"username"`
	Icon              string              `json:"icon"`
	Bio               string              `json:"comment"`
	Sns               string              `json:"SNS"`
}

type SubmissionJson struct{
	ID                int                 `json:"submission_id"`
	RequestID         int                 `json:"request_id"`
	CreatedAt         time.Time           `json:"createdat"`
	EngineerProfile   UserProfileJson     `json:"engineer"`
	URL               string              `json:url`
	Content           string              `json:"content"`
}

type RequestJson struct{
	ID                int                 `json:"request_id"`
	RequestName       string              `json:"requestname"`
	CreatedAt         time.Time           `json:"createdat"`
	ClientProfile     UserProfileJson     `json:"client"`
	EngineersProfile  []UserProfileJson   `json:"engineers"`
	Content           string              `json:"content"`
	Submissions       []SubmissionJson    `json:"submissions"`
}

type RequestsJson struct{
	Requests          []RequestJson       `json:"requests"`
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
		// Client構造体を格納するためのインスタンスを生成
		client := typefile.Client{}
		// ClientのProfile構造体を格納するためのインスタンスを生成
		client_profile := typefile.Profile{}
		// Submission構造体データを複数格納するためのインスタンスを生成
		submissions := []typefile.Submission{}
		// Engineer構造体を複数格納するためのインスタンスを生成
		engineers := []typefile.Engineer{}

		//　特定リクエストのidを持つsubmissionを格納する。
		model.Db.Find(&submissions,"request_id = ?",request.ID)
		// SELECT * FROM `submissions` WHERE request_id = ?
		// 特定のidを持つclientを抽出する。
		model.Db.Find(&client,"id = ?",request.ClientID)
		// SELECT * FROM `clients` WHERE id = ?
		// 特定のidを持つprofileを抽出する。
		model.Db.Find(&client_profile,"id = ?",request.ClientID)
		// SELECT * FROM `profiles` WHERE id = ?
		// Associationによって、engineerデータを取り出す。
		model.Db.Model(&request).Association("Engineers").Find(&engineers)

		// requestが持つデータをrequest_jsonのそれぞれの対応する属性に格納する。
		request_json.ID = request.ID
		request_json.RequestName = request.RequestName
		request_json.CreatedAt = request.CreatedAt
		request_json.Content = request.Content
		request_json.ClientProfile.UserID = request.ClientID
		request_json.ClientProfile.UserName = client.User.Name
		request_json.ClientProfile.Icon = client_profile.Icon
		request_json.ClientProfile.Bio = client_profile.Bio
		request_json.ClientProfile.Sns = string(client_profile.Sns)
		request_json.Content = request.Content
		// EngineersIDは複数存在するため、エンジニアのIDをループで追加していく。
		for _, engineer := range engineers{
			// EngineerのUserProfileJson構造体を格納するためのインスタンスを生成
			engineer_profile_json := UserProfileJson{}
			// EngineerのProfile構造体を格納するためのインスタンスを生成
			engineer_profile := typefile.Profile{}

			// 特定のidを持つprofileを抽出する。
			model.Db.Find(&engineer_profile,"id = ?",engineer.User.ID)
			// SELECT * FROM `profiles` WHERE id = ?

			// requestが持つデータをrequest_jsonのそれぞれの対応する属性に格納する。
			engineer_profile_json.UserID = engineer.User.ID
			engineer_profile_json.UserName = engineer.User.Name
			engineer_profile_json.Icon = engineer_profile.Icon
			engineer_profile_json.Bio = engineer_profile.Bio
			engineer_profile_json.Sns = string(engineer_profile.Sns)

			request_json.EngineersProfile = append(request_json.EngineersProfile,engineer_profile_json)
		}
		// submissionは複数存在するため、submissionデータをループで追加していく。
		for _, submission := range submissions{
			// SubmissionJson構造体データを格納するためのインスタンスを生成
			submission_json := SubmissionJson{}
			// Engineer構造体を格納するためのインスタンスを生成
			engineer := typefile.Engineer{}
			// EngineerのProfile構造体を格納するためのインスタンスを生成
			engineer_profile := typefile.Profile{}

			// 特定のidを持つengineerを抽出する。
			model.Db.Find(&engineer,"id = ?",submission.EngineerID)
			// SELECT * FROM `engineers` WHERE id = ?
			// 特定のidを持つprofileを抽出する。
			model.Db.Find(&engineer_profile,"id = ?",submission.EngineerID)
			// SELECT * FROM `profiles` WHERE id = ?

			// submissionが持つデータをsubmission_jsonのそれぞれの対応する属性に格納する。
			submission_json.ID = submission.ID
			submission_json.RequestID = submission.RequestID
			submission_json.CreatedAt = submission.CreatedAt
			submission_json.URL = submission.URL
			submission_json.EngineerProfile.UserID = engineer.User.ID
			submission_json.EngineerProfile.UserName = engineer.User.Name
			submission_json.EngineerProfile.Icon = engineer_profile.Icon
			submission_json.EngineerProfile.Bio = engineer_profile.Bio
			submission_json.EngineerProfile.Sns = string(engineer_profile.Sns)
			submission_json.Content = submission.Content

			request_json.Submissions = append(request_json.Submissions,submission_json)
		}
		requests_json.Requests = append(requests_json.Requests,request_json)
	}

	c.JSON(http.StatusOK, gin.H{
    	"requests": requests_json.Requests,
    })
}