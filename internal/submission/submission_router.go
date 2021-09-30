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
type SubmissionJson struct{
	ID             int                 `json:"submission_id"`
	CreatedAt      time.Time           `json:"createdat"`
	// 要件はエンジニアのプロフィールデータであるが、プロフィール機能は担当外のため、EngineerIDを代用する。
	EngineerID     int                 `json:"engineer`
	Content        string              `json:"content"`
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
	
	// 特定のidを持つsubmissionを抽出する。
	model.Db.Find(&submission,"id = ?",submission_id)
	// SELECT * FROM `submissions` WHERE id = ?

	// submissionが持つデータをsubmission_jsonのそれぞれの対応する属性に格納する。
	submission_json.ID = submission.ID
	submission_json.CreatedAt = submission.CreatedAt
	submission_json.EngineerID = submission.EngineerID
	submission_json.Content = submission.Content

    c.JSON(http.StatusOK, gin.H{
    	"submission": submission_json,
    })
}