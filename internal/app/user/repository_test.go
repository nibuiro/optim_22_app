package user

import (
  "context"
  "testing"
  "regexp"
  "github.com/DATA-DOG/go-sqlmock"
  "github.com/stretchr/testify/suite"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "optim_22_app/pkg/log"
  "optim_22_app/typefile"

)

/*
 *  mysqlドライバに適合した書き方が不明。
 *  postgresドライバで代用
 *
 *  参考：
 *　[INSERT while mocking gorm · Issue #118 · DATA-DOG/go-sqlmock · GitHub]
 *  (https://github.com/DATA-DOG/go-sqlmock/issues/118)
 *
 */


// テストスイートの構造体
type UserRepositoryTestSuite struct {
  suite.Suite
  userRepository repository
  mock           sqlmock.Sqlmock
}

// テストのセットアップ
// (sqlmockをNew、Gormで発行されるクエリがモックに送られるように)
func (suite *UserRepositoryTestSuite) SetupTest() {
  db, mock, _ := sqlmock.New()
  suite.mock = mock
  logger := log.New()
  DB, _ := gorm.Open(
    postgres.New(postgres.Config{Conn: db,}), 
    &gorm.Config{},
  )
  userRepository := repository{DB, logger}
  suite.userRepository = userRepository
}

// テスト終了時の処理（データベース接続のクローズ）
func (suite *UserRepositoryTestSuite) TearDownTest() {
  db, _ := suite.userRepository.db.DB()
  db.Close()
}

// テストスイートの実行
func TestUserRepositoryTestSuite(t *testing.T) {
  suite.Run(t, new(UserRepositoryTestSuite))
}

// Createのテスト
func (suite *UserRepositoryTestSuite) TestCreate() {
  suite.Run("create a user", func() {
      newId := 1
      rows := sqlmock.NewRows([]string{"id"}).AddRow(newId)
      suite.mock.ExpectBegin()
      suite.mock.ExpectQuery(
        regexp.QuoteMeta(`INSERT INTO "users" ("name","email","password") VALUES ($1,$2,$3) RETURNING "id"`),
      ).
      WillReturnRows(rows)
      suite.mock.ExpectCommit()
      
      insertValues := &typefile.User{
        Name:      "test",
        Email:     "test@test.test",
        Password:  "test",
      }
      
      var ctx context.Context
      err := suite.userRepository.Create(ctx, insertValues)
    
      suite.userRepository.logger.Debug(insertValues)
      suite.userRepository.logger.Debug(err)
    },
  )
}
