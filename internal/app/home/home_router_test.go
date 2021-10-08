package home

import (
	"optim_22_app/internal/app/home"
	"optim_22_app/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)


// 指定のパスが入力された時に、ステータスコード200を返すかを確認している。


// 特定リクエストの詳細を表示する関数のテスト。200ステータスコードを返すテストを実行している。
func TestSuccesshomeShowHomepage(t *testing.T) {
	router := gin.New()

	router.GET("/api/requests",home.ShowHomepage)

	model.InitDB()

	w := httptest.NewRecorder()
	url := "/api/requests"
	req, _ := http.NewRequest("GET", url, nil)
	// HandlerFuncに対して、ServeHTTP(w ResponceWriter,r *Request)を使うことで、
	// 実際にサーバーを立ち上げずにリクエストをシミュレートすることができる。
	router.ServeHTTP(w,req)

	assert.Equal(t, http.StatusOK, w.Code)
}