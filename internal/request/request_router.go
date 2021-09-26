package request

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"optim_22_app/model"
	"typefile"
	"strconv"
)

// 特定リクエストの詳細を表示する
func ShowRequest(c *gin.Context) {
	// urlの引数で受け取ったengineer_idをengineer_id_stringという変数に格納している。
	request_id_string := c.Param("request_id")
	// 文字列をintに変換
	request_id, _ := strconv.Atoi(request_id_string)

	engineer_id := 3

	// Request構造体を格納するための型を宣言
	request := typefile.Request{}
	// Winner構造体を格納するための型を宣言
	winner := typefile.Winner{}
	// Engineer構造体を格納するための型を宣言
	winner_engineer := typefile.Engineer{}
	// Client構造体を格納するための型を宣言
	client := typefile.Client{}
	
	// 特定のidを持つRequestを抽出する。
	model.Db.Find(&request,"id = ?",request_id)
	// SELECT * FROM `requests` WHERE id = ?
	// 特定のrequest_idを持つwinnerを抽出する。
	model.Db.Find(&winner,"request_id = ?",request_id)
	// SELECT * FROM `winners` WHERE request_id = ?
	// 特定のidを持つclientを抽出する。
	model.Db.Find(&client,"id = ?",request.ClientID)
	// SELECT * FROM `clients` WHERE id = ?

	// winnerがいたか確認するための変数。初期値はfalseとする。
	winner_is_empty := false

	// winnerがいるか確認。空の場合、request_is_emptyにtrueを渡す。空でない場合、falseを渡す。
	// 特定のリクエストidを持つwinnerを抽出しなかった場合、データが存在しないため、エンジニアIDは0を返す。
	if winner.EngineerID == 0{
		winner_is_empty = true
	}else{
		winner_is_empty = false
		// winnerであるエンジニアを抽出する。
		model.Db.Find(&winner_engineer,"id = ?",winner.EngineerID)
	}

    c.HTML(http.StatusOK, "show_request_detail.html", gin.H{
    	"winner_is_empty": winner_is_empty,
    	"winner_engineer": winner_engineer,
    	"client": client,
    	"request": request,
    	"engineer_id": engineer_id,
    })
}