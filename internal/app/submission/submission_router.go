package submission

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"optim_22_app/model"
	"typefile"
	"strconv"
	"time"
)

//データを送信する際に利用する構造体を定義
type UserProfileJson struct{
	UserID            int                   `json:"user_id"`
	UserName          string                `json:"username"`
	Icon              string                `json:"icon"`
	Bio               string                `json:"comment"`
	Sns               string                `json:"SNS"`
}

type SubmissionJson struct{
	ID                int                   `json:"submission_id"`
	RequestID         int                   `json:"request_id"`
	CreatedAt         time.Time             `json:"createdat"`
	EngineerProfile   UserProfileJson       `json:"engineer"`
	URL               string                `json:url`
	Content           string                `json:"content"`
}

// 特定submissionの詳細を表示する
func ShowSubmission(c *gin.Context) {
	// urlの引数で受け取ったsubmission_idをsubmission_id_stringという変数に格納している。
	submission_id_string := c.Param("submission_id")
	// 文字列をintに変換
	submission_id, _ := strconv.Atoi(submission_id_string)

	// Submission構造体データを格納するためのインスタンスを生成
	submission := typefile.Submission{}
	// SubmissionJson構造体データを格納するためのインスタンスを生成
	submission_json := SubmissionJson{}
	// Engineer構造体を格納するためのインスタンスを生成
	engineer := typefile.Engineer{}
	// EngineerのProfile構造体を格納するためのインスタンスを生成
	engineer_profile := typefile.Profile{}
	
	// 特定のidを持つsubmissionを抽出する。
	model.Db.Find(&submission,"id = ?",submission_id)
	// SELECT * FROM `submissions` WHERE id = ?
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
	submission_json.Content = submission.Content
	submission_json.EngineerProfile.UserID = engineer.User.ID
	submission_json.EngineerProfile.UserName = engineer.User.Name
	submission_json.EngineerProfile.Icon = engineer_profile.Icon
	submission_json.EngineerProfile.Bio = engineer_profile.Bio
	submission_json.EngineerProfile.Sns = string(engineer_profile.Sns)

	if submission.ID == 0{
		// 400ラーを返したいが指定方法が分からない。なので、存在しないファイルを指定することで、404errorを出させる。
		c.JSON(http.StatusNotFound, gin.H{})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"submission": submission_json,
		})
	}
}