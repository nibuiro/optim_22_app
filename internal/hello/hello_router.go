package hello

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Hello(c *gin.Context) {
    c.HTML(http.StatusOK, "hello.html", gin.H{})
}

func NewHello(c *gin.Context) {
	// c.QueryでURLクエリ文字列を取得できる。取得できる値の型はstringである。
	// https://github.com/gin-gonic/gin#querystring-parametersに詳細が記載されている。
	request_id_string := c.Query("request_id")
	engineer_id_string := c.Query("engineer_id")
	test_id_string := c.Query("test_id")

	// 文字列をintに変換
	request_id, _ := strconv.Atoi(request_id_string)
	engineer_id, _ := strconv.Atoi(engineer_id_string)
	test_id, _ := strconv.Atoi(test_id_string)
	c.HTML(http.StatusOK, "newhello.html", gin.H{
		"request_id": request_id,
		"engineer_id": engineer_id,
		"test_id": test_id,
	})
}

