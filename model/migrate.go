// https://gorm.io/docs/connecting_to_the_database.html#Clickhouse に詳細が記載されている。

package model

import (
	"fmt"
	"typefile"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

	// マイグレーションは定義したstructをAutoMigrateの引数に渡すことで、
	// それに対応するテーブルの作成を行う。
	// テーブル作成時にオプションを付けたい場合、db.Set()を利用する。
	db.AutoMigrate(&typefile.Request{})

	// Insert
	// db.Create(&request)

	// Select
    // db.Find(&request, "id = ?", 10)

    // Batch Insert
    // var requests = []User{request1, request2, request3}
    // db.Create(&users)
}