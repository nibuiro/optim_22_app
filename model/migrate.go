// maikl"os"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"typefile"
)

// 外部でdb操作をするためのパッケージ変数
var Db *gorm.DB

func InitDB() {
	var err error
	// maikl"root:rootpass@tcp(mysql_container:3306)/optim_dev?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"
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

	icon, err := os.ReadFile("/go/src/optim_22_app/model/icon")

	fmt.Println(err)

	var users = []typefile.User{
		{Name: "user1", Email: "test@test.test", Password: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"},
		{Name: "user2", Email: "test@inc.test-ac.jp", Password: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"},
		{Name: "user3", Email: "test.test@test.com", Password: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"},
		{Name: "user4", Email: "test@test.test", Password: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"},
		{Name: "user5", Email: "test@inc.test-ac.jp", Password: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"}}
	Db.Create(&users)

	var profile = []typefile.Profile{
		{ID: 1,Bio: "C#", Sns: []byte(`{"github": "pole", "twitter": "pole", "facebook": "pole"}`), Icon: string(icon)},
		{ID: 2,Bio: "C++", Sns: []byte(`{"github": "maikl", "twitter": "maikl", "facebook": "maikl"}`), Icon: string(icon)},
		{ID: 3,Bio: "C", Sns: []byte(`{"github": "suzuki", "twitter": "suzuki", "facebook": "suzuki"}`), Icon: string(icon)},
		{ID: 4,Bio: "haskell", Sns: []byte(`{"github": "siraisi", "twitter": "siraisi", "facebook": "siraisi"}`), Icon: string(icon)},
		{ID: 5,Bio: "golang", Sns: []byte(`{"github": "tom", "twitter": "tom", "facebook": "tom"}`), Icon: string(icon)},
	}
	Db.Create(&profile)

	// userが作成された直後にengineerも作成する。
	var engineers = []typefile.Engineer{
		{User: typefile.User{ID: 1,Name: "user1", Email: "test@test.test"}},
		{User: typefile.User{ID: 2,Name: "user2", Email: "test@inc.test-ac.jp"}},
		{User: typefile.User{ID: 3,Name: "user3", Email: "test.test@test.com"}},
		{User: typefile.User{ID: 4,Name: "user4", Email: "test@test.test"}},
		{User: typefile.User{ID: 5,Name: "user5", Email: "test@inc.test-ac.jp"}}}
	Db.Create(&engineers)

	//userが作成された直後にclientも作成する。
	var clients = []typefile.Client{
		{User: typefile.User{ID: 1,Name: "user1", Email: "test@test.test"}},
		{User: typefile.User{ID: 2,Name: "user2", Email: "test@inc.test-ac.jp"}},
		{User: typefile.User{ID: 3,Name: "user3", Email: "test.test@test.com"}},
		{User: typefile.User{ID: 4,Name: "user4", Email: "test@test.test"}},
		{User: typefile.User{ID: 5,Name: "user5", Email: "test@inc.test-ac.jp"}}}
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
		{RequestID: 1,EngineerID: 2,URL: "http://example.com/3",Content: "submission3 of engineerID 2"},
		{RequestID: 4,EngineerID: 2,URL: "http://example.com/4",Content: "submission4 of engineerID 2"},
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
		engineer2}

	// リクエスト1に参加しているエンジニアを外部キーなしで取得するために、Associationを追加している。
	Db.Model(&request1).Association("Engineers").Append(&request1_engineers_association)

	var request2_engineers_association = []typefile.Engineer{
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
		request3}

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

	var comments = []typefile.Comment{
		{UserID: 1,RequestID: 1,Title: "サンプルコメントタイトル1",Body: "サンプルコメント本文1",ReplyID: 0,},
		{UserID: 2,RequestID: 1,Title: "サンプルコメントタイトル2",Body: "サンプルコメント本文2",ReplyID: 1,},
		{UserID: 1,RequestID: 1,Title: "サンプルコメントタイトル3",Body: "サンプルコメント本文3",ReplyID: 2,},
		{UserID: 3,RequestID: 1,Title: "サンプルコメントタイトル4",Body: "サンプルコメント本文4",ReplyID: 3,},
		{UserID: 1,RequestID: 2,Title: "サンプルコメントタイトル5",Body: "サンプルコメント本文5",ReplyID: 0,},
	}
	Db.Create(&comments)
}