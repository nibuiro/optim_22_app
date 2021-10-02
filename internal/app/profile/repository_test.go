package profile

import (
  "context"
  "testing"
  "regexp"
  "github.com/DATA-DOG/go-sqlmock"
  "github.com/stretchr/testify/suite"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "optim_22_app/pkg/log"
//  "optim_22_app/typefile"

)


// テストスイートの構造体
type ProfileRepositoryTestSuite struct {
  suite.Suite
  repository repository
  mock           sqlmock.Sqlmock
}

// テストのセットアップ
// (sqlmockをNew、Gormで発行されるクエリがモックに送られるように)
func (suite *ProfileRepositoryTestSuite) SetupTest() {
  db, mock, _ := sqlmock.New()
  suite.mock = mock
  logger := log.New()
  DB, _ := gorm.Open(
    postgres.New(postgres.Config{Conn: db,}), 
    &gorm.Config{},
  )
  repository := repository{DB, logger}
  suite.repository = repository
}

// テスト終了時の処理（データベース接続のクローズ）
func (suite *ProfileRepositoryTestSuite) TearDownTest() {
  db, _ := suite.repository.db.DB()
  db.Close()
}

// テストスイートの実行
func TestProfileRepositoryTestSuite(t *testing.T) {
  suite.Run(t, new(ProfileRepositoryTestSuite))
}

// Createのテスト
func (suite *ProfileRepositoryTestSuite) TestCreate() {
  suite.Run("create a entity.request, not a http request.", func() {
      newId := 1
      rows := sqlmock.NewRows([]string{})//.AddRow(newId)
      //suite.mock.ExpectBegin()
      suite.mock.ExpectQuery(
        regexp.QuoteMeta(`SELECT request.Finish, submission.UpdatedAt, request.ClientID, request.RequestName, request.Content, submissionID FROM "submissions" INNER JOIN "request" ON submission.RequestID = request.ID WHERE submission.EngineerID = $1`),
      ).
      WithArgs(1).
      WillReturnRows(rows)
      //suite.mock.ExpectCommit()
      
      var ctx context.Context
      _, err := suite.repository.GetParticipated(ctx, newId)
    
      //suite.repository.logger.Debug(insertValues)
      suite.repository.logger.Debug(err)
    },
  )
}