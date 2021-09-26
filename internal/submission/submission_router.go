package submission

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"optim_22_app/model"
	"typefile"
	"strconv"
)

// 特定submissionの詳細を表示する
func ShowSubmission(c *gin.Context) {
	// urlの引数で受け取ったsubmission_idをsubmission_id_stringという変数に格納している。
	submission_id_string := c.Param("submission_id")
	// 文字列をintに変換
	submission_id, _ := strconv.Atoi(submission_id_string)

	// Submission構造体データを格納するためのインスタンスを生成
	submission := typefile.Submission{}
	// Request構造体データを格納するためのインスタンスを生成
	request := typefile.Request{}
	
	// 特定のidを持つsubmissionを抽出する。
	model.Db.Find(&submission,"id = ?",submission_id)
	// SELECT * FROM `submissions` WHERE id = ?
	// 特定のidを持つrequestを抽出する。
	model.Db.Find(&request,"id = ?",submission.RequestID)
	// SELECT * FROM `requests` WHERE id = ?

    c.HTML(http.StatusOK, "show_submission_detail.html", gin.H{
    	"submission": submission,
    	"request": request,
    })
}