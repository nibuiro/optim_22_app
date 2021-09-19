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
type RequestRepositoryTestSuite struct {
  suite.Suite
  requestRepository requestRepository
  mock           sqlmock.Sqlmock
}

// テストのセットアップ
// (sqlmockをNew、Gormで発行されるクエリがモックに送られるように)
func (suite *RequestRepositoryTestSuite) SetupTest() {
  db, mock, _ := sqlmock.New()
  suite.mock = mock
  logger := log.New()
  DB, _ := gorm.Open(
    postgres.New(postgres.Config{Conn: db,}), 
    &gorm.Config{},
  )
  requestRepository := requestRepository{DB, logger}
  suite.requestRepository = requestRepository
}

// テスト終了時の処理（データベース接続のクローズ）
func (suite *RequestRepositoryTestSuite) TearDownTest() {
  db, _ := suite.requestRepository.db.DB()
  db.Close()
}

// テストスイートの実行
func TestRequestRepositoryTestSuite(t *testing.T) {
  suite.Run(t, new(RequestRepositoryTestSuite))
}

// Createのテスト
func (suite *RequestRepositoryTestSuite) TestCreate() {
  suite.Run("create a entity.request, not a http request.", func() {
      newId := 1
      rows := sqlmock.NewRows([]string{})//.AddRow(newId)
      //suite.mock.ExpectBegin()
      suite.mock.ExpectQuery(
        regexp.QuoteMeta(`SELECT request.Finish, submission.UpdatedAt, request.ClientID, request.RequestName, request.Content, submissionID FROM "submissions" left join "request" on submission.RequestID = request.ID WHERE submission.EngineerID = $1`),
      ).
      WithArgs(1).
      WillReturnRows(rows)
      //suite.mock.ExpectCommit()
      
      var ctx context.Context
      _, err := suite.requestRepository.GetParticipated(ctx, newId)
    
      //suite.requestRepository.logger.Debug(insertValues)
      suite.requestRepository.logger.Debug(err)
    },
  )
}