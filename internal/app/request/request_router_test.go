package request

import (
	"optim_22_app/internal/app/request"
	"optim_22_app/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)


// 存在するrequestデータのidが渡されたとき、存在しないrequestデータのidが渡されたときのステータスコードを確認している。


// 特定リクエストの詳細を表示する関数のテスト。200ステータスコードを返すテストを実行している。
func TestSuccessShowRequest(t *testing.T) {
	router := gin.New()

	router.GET("request/show_request/:request_id",request.ShowRequest)

	model.InitDB()

	// テストするrequest_idを格納している。
	test_ids := []int{1,2,3}
	for _, test_id := range test_ids{
		w := httptest.NewRecorder()
		url := "/request/show_request/" + strconv.Itoa(test_id)
		req, _ := http.NewRequest("GET", url, nil)
		// HandlerFuncに対して、ServeHTTP(w ResponceWriter,r *Request)を使うことで、
		// 実際にサーバーを立ち上げずにリクエストをシミュレートすることができる。
		router.ServeHTTP(w,req)

		assert.Equal(t, http.StatusOK, w.Code)
	}
}

// 特定リクエストの詳細を表示する関数のテスト。400ステータスコードを返すテストを実行している。
func TestFailureShowRequest(t *testing.T) {
	router := gin.New()

	router.GET("request/show_request/:request_id",request.ShowRequest)

	model.InitDB()

	// テストするrequest_idを格納している。
	test_ids := []int{1001,1002,1003}
	for _, test_id := range test_ids{
		w := httptest.NewRecorder()
		url := "/request/show_request/" + strconv.Itoa(test_id)
		req, _ := http.NewRequest("GET", url, nil)
		// HandlerFuncに対して、ServeHTTP(w ResponceWriter,r *Request)を使うことで、
		// 実際にサーバーを立ち上げずにリクエストをシミュレートすることができる。
		router.ServeHTTP(w,req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	}
}