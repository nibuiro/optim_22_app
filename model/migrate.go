// https://gorm.io/docs/connecting_to_the_database.html#Clickhouse に詳細が記載されている。

package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"typefile"
)

// 外部でdb操作をするためのパッケージ変数
var Db *gorm.DB

func InitDB() {
	var err error
	// https://github.com/go-sql-driver/mysql#dsn-data-source-name に詳細が記載されている。
	// DSN(データソース名)の作成。
	// 開発用のデータベース名はoptim_dev,テスト用のデータベース名はotpim_testである。
	dsn := "root:rootpass@tcp(mysql_container:3306)/optim_dev"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("database successfully configure")
	}

	// 接続したdbをパッケージ変数Dbに代入している。
	Db = db
}

// Insert
// db.Create(&request)

// Select
// db.Find(&request, "id = ?", 10)

// Batch Insert
// var requests = []User{request1, request2, request3}
// db.Create(&requests)


// テストを実行するために前もって必要なデータを作成する。
func CreateTestData() {
	var users = []typefile.User{
		{Name: "user1"},
		{Name: "user2"},
		{Name: "user3"},
		{Name: "user4"},
		{Name: "user5"},
		{Name: "user6"}}
	Db.Create(&users)

	var engineers = []typefile.Engineer{
		{User: typefile.User{ID: 1,Name: "user1"}},
		{User: typefile.User{ID: 2,Name: "user2"}},
		{User: typefile.User{ID: 3,Name: "user3"}}}
	Db.Create(&engineers)
	
	var clients = []typefile.Client{
		{User: typefile.User{ID: 4,Name: "user4"}},
		{User: typefile.User{ID: 5,Name: "user5"}},
		{User: typefile.User{ID: 6,Name: "user6"}}}
	Db.Create(&clients)

	var requests = []typefile.Request{
		{ClientID: 4,RequestName: "request1",Content: "request1 content",Finish: true},
		{ClientID: 4,RequestName: "request2",Content: "request2 content",Finish: true},
		{ClientID: 5,RequestName: "request3",Content: "request3 content",Finish: false}}
	Db.Create(&requests)
	
	var winners = []typefile.Winner{
		{User: typefile.User{ID: 1,Name: "user1"},RequestID: 1},
		{User: typefile.User{ID: 2,Name: "user2"},RequestID: 2}}
	Db.Create(&winners)
}