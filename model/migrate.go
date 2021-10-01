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
	dsn := "root:rootpass@tcp(mysql_container:3306)/optim_dev?parseTime=true&loc=Asia%2FTokyo"
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
		{Name: "user5"}}
	Db.Create(&users)

	// userが作成された直後にengineerも作成する。
	var engineers = []typefile.Engineer{
		{User: typefile.User{ID: 1,Name: "user1"}},
		{User: typefile.User{ID: 2,Name: "user2"}},
		{User: typefile.User{ID: 3,Name: "user3"}},
		{User: typefile.User{ID: 4,Name: "user4"}},
		{User: typefile.User{ID: 5,Name: "user5"}}}
	Db.Create(&engineers)

	//userが作成された直後にclientも作成する。
	var clients = []typefile.Client{
		{User: typefile.User{ID: 1,Name: "user1"}},
		{User: typefile.User{ID: 2,Name: "user2"}},
		{User: typefile.User{ID: 3,Name: "user3"}},
		{User: typefile.User{ID: 4,Name: "user4"}},
		{User: typefile.User{ID: 5,Name: "user5"}}}
	Db.Create(&clients)

	var requests = []typefile.Request{
		{ClientID: 1,RequestName: "request1 from clientID 1",Content: "request1 content",Finish: false},
		{ClientID: 1,RequestName: "request2 from clientID 1",Content: "request2 content",Finish: false},
		{ClientID: 2,RequestName: "request3 from clientID 2",Content: "request3 content",Finish: true},
		{ClientID: 3,RequestName: "request4 from clientID 1",Content: "request4 content",Finish: false},
		{ClientID: 3,RequestName: "request5 from clientID 1",Content: "request5 content",Finish: true}}
	Db.Create(&requests)
	
	var winners = []typefile.Winner{
		{EngineerID: 1,RequestID: 3},
		{EngineerID: 4,RequestID: 5}}
	Db.Create(&winners)

	var submissions = []typefile.Submission{
		{RequestID: 1,EngineerID: 1,URL: "http://example.com/1",Content: "submission1 of engineerID 1"},
		{RequestID: 2,EngineerID: 1,URL: "http://example.com/2",Content: "submission2 of engineerID 1"},
		{RequestID: 1,EngineerID: 2,URL: "http://example.com/3",Content: "submission3 of engineerID 2"},
		{RequestID: 4,EngineerID: 2,URL: "http://example.com/4",Content: "submission4 of engineerID 2"},
		{RequestID: 4,EngineerID: 3,URL: "http://example.com/5",Content: "submission5 of engineerID 3"},
		{RequestID: 3,EngineerID: 1,URL: "http://example.com/6",Content: "submission6 of engineerID 1"},
		{RequestID: 5,EngineerID: 4,URL: "http://example.com/7",Content: "submission7 of engineerID 4"}}
	Db.Create(&submissions)

	// id=1のRequest構造体データを格納するためのインスタンスを生成
	request1 := typefile.Request{}
	// id=1を持つrequestを抽出する。
	Db.Find(&request1,"id = ?",1)
	// SELECT * FROM `requests` WHERE id = 1

	// id=2のRequest構造体データを格納するためのインスタンスを生成
	request2 := typefile.Request{}
	// id=2を持つrequestを抽出する。
	Db.Find(&request2,"id = ?",2)
	// SELECT * FROM `requests` WHERE id = 2

	// id=3のRequest構造体データを格納するためのインスタンスを生成
	request3 := typefile.Request{}
	// id=3を持つrequestを抽出する。
	Db.Find(&request3,"id = ?",3)
	// SELECT * FROM `requests` WHERE id = 3

	// id=4のRequest構造体データを格納するためのインスタンスを生成
	request4 := typefile.Request{}
	// id=4を持つrequestを抽出する。
	Db.Find(&request4,"id = ?",4)
	// SELECT * FROM `requests` WHERE id = 4

	// id=5のRequest構造体データを格納するためのインスタンスを生成
	request5 := typefile.Request{}
	// id=5を持つrequestを抽出する。
	Db.Find(&request5,"id = ?",5)
	// SELECT * FROM `requests` WHERE id = 5

	// id=1のEngineer構造体データを格納するためのインスタンスを生成
	engineer1 := typefile.Engineer{}
	// id=1を持つengineerを抽出する。
	Db.Find(&engineer1,"id = ?",1)
	// SELECT * FROM `engineers` WHERE id = 1

	// id=2のEngineer構造体データを格納するためのインスタンスを生成
	engineer2 := typefile.Engineer{}
	// id=2を持つengineerを抽出する。
	Db.Find(&engineer2,"id = ?",2)
	// SELECT * FROM `engineers` WHERE id = 2

	// id=3のEngineer構造体データを格納するためのインスタンスを生成
	engineer3 := typefile.Engineer{}
	// id=3を持つengineerを抽出する。
	Db.Find(&engineer3,"id = ?",3)
	// SELECT * FROM `engineers` WHERE id = 3

	// id=4のEngineer構造体データを格納するためのインスタンスを生成
	engineer4 := typefile.Engineer{}
	// id=4を持つengineerを抽出する。
	Db.Find(&engineer4,"id = ?",4)
	// SELECT * FROM `engineers` WHERE id = 4

	// id=5のEngineer構造体データを格納するためのインスタンスを生成
	engineer5 := typefile.Engineer{}
	// id=5を持つengineerを抽出する。
	Db.Find(&engineer5,"id = ?",5)
	// SELECT * FROM `engineers` WHERE id = 5

	var request1_engineers_association = []typefile.Engineer{
		engineer1,
		engineer2}

	// リクエスト1に参加しているエンジニアを外部キーなしで取得するために、Associationを追加している。
	Db.Model(&request1).Association("Engineers").Append(&request1_engineers_association)

	var request2_engineers_association = []typefile.Engineer{
		engineer1,
		engineer2}

	// リクエスト2に参加しているエンジニアを外部キーなしで取得するために、Associationを追加している。
	Db.Model(&request2).Association("Engineers").Append(&request2_engineers_association)

	var request3_engineers_association = []typefile.Engineer{
		engineer1,
		engineer3,
		engineer5}

	// リクエスト3に参加しているエンジニアを外部キーなしで取得するために、Associationを追加している。
	Db.Model(&request3).Association("Engineers").Append(&request3_engineers_association)

	var request4_engineers_association = []typefile.Engineer{
		engineer2,
		engineer3,
		engineer4,
		engineer5}

	// リクエスト4に参加しているエンジニアを外部キーなしで取得するために、Associationを追加している。
	Db.Model(&request4).Association("Engineers").Append(&request4_engineers_association)

	var request5_engineers_association = []typefile.Engineer{
		engineer4,
		engineer5}
	
	// リクエスト5に参加しているエンジニアを外部キーなしで取得するために、Associationを追加している。
	Db.Model(&request5).Association("Engineers").Append(&request5_engineers_association)

	var engineer1_requests_association = []typefile.Request{
		request1,
		request2,
		request3}

	// エンジニア1が参加しているリクエストを外部キーなしで取得するために、Associationを追加している。
	Db.Model(&engineer1).Association("Requests").Append(&engineer1_requests_association)

	var engineer2_requests_association = []typefile.Request{
		request1,
		request2,
		request4}

	// エンジニア2が参加しているリクエストを外部キーなしで取得するために、Associationを追加している。
	Db.Model(&engineer2).Association("Requests").Append(&engineer2_requests_association)

	var engineer3_requests_association = []typefile.Request{
		request3,
		request4}

	// エンジニア3が参加しているリクエストを外部キーなしで取得するために、Associationを追加している。
	Db.Model(&engineer3).Association("Requests").Append(&engineer3_requests_association)

	var engineer4_requests_association = []typefile.Request{
		request4,
		request5}

	// エンジニア4が参加しているリクエストを外部キーなしで取得するために、Associationを追加している。
	Db.Model(&engineer4).Association("Requests").Append(&engineer4_requests_association)

	var engineer5_requests_association = []typefile.Request{
		request3,
		request4,
		request5}

	// エンジニア5が参加しているリクエストを外部キーなしで取得するために、Associationを追加している。
	Db.Model(&engineer5).Association("Requests").Append(&engineer5_requests_association)
}