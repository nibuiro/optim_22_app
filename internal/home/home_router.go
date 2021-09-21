package home

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"optim_22_app/model"
	"typefile"
)

// ホームページを表示する
func ShowHomepage(c *gin.Context) {
	// Request構造体データを複数格納するためのインスタンスを生成
	requests := []typefile.Request{}

	//　受付中のリクエストを格納する。(Winnerを持たないリクエストを抽出する)
	model.Db.Find(&requests,"finish = ?",0)
	// SELECT * FROM `request` WHERE finish = ?

	// Client構造体データを複数格納するためのインスタンスを生成
	clients := []typefile.Client{}
	// Engineer構造体データを複数格納するためのインスタンスを生成
	engineers := []typefile.Engineer{}

	// Associationによって、clientデータを取り出す。
	model.Db.Model(&requests).Association("Client").Find(&clients)
	// Associationによって、engineerデータを取り出す。
	model.Db.Model(&requests).Association("Engineer").Find(&engineers)

    c.HTML(http.StatusOK, "show_homepage.html", gin.H{
    	"requests": requests,
    	"clients": clients,
    	"engineers": engineers,
    })
}