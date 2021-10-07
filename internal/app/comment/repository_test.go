package comment

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


// テストスイートの構造体
type CommentRepositoryTestSuite struct {
  suite.Suite
  repository repository
  mock           sqlmock.Sqlmock
}

// テストのセットアップ
// (sqlmockをNew、Gormで発行されるクエリがモックに送られるように)
func (suite *CommentRepositoryTestSuite) SetupTest() {
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
func (suite *CommentRepositoryTestSuite) TearDownTest() {
  db, _ := suite.repository.db.DB()
  db.Close()
}

// テストスイートの実行
func TestCommentRepositoryTestSuite(t *testing.T) {
  suite.Run(t, new(CommentRepositoryTestSuite))
}


func (suite *CommentRepositoryTestSuite) TestCreate() {

  suite.Run("post a comment", func() {
      newId := 1
      rows := sqlmock.NewRows([]string{"id"}).AddRow(newId)
      suite.mock.ExpectBegin()
      suite.mock.ExpectQuery(
        regexp.QuoteMeta(`INSERT INTO "comments" ("created_at","request_id","user_id","title","body","reply_id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`),
      ).
      WillReturnRows(rows)
      suite.mock.ExpectCommit()

      insertValues := &typefile.Comment{
        RequestID: 1,
        UserID: 1,
        Title: "test",
        Body: "test",
        ReplyID: 0,
      }
      
      var ctx context.Context
      err := suite.repository.Create(ctx, insertValues)
    
      suite.repository.logger.Debug(insertValues)
      suite.repository.logger.Debug(err)
    },
  )

  suite.Run("get comments", func() {
      requestID := 1
      rows := sqlmock.NewRows([]string{"id"}).AddRow(requestID)
      //suite.mock.ExpectBegin()
      suite.mock.ExpectQuery(
        regexp.QuoteMeta(`SELECT comments.ID, comments.RequestID, comments.UserID, user.Name, comments.CreatedAt, comments.Title, comments.Body, comments.ReplyID FROM "comments" INNER JOIN "user" ON comments.userID = user.ID WHERE comments.RequestID = $1`),
      ).
      WillReturnRows(rows)

      var ctx context.Context
      _, err := suite.repository.Get(ctx, requestID)
    
      suite.repository.logger.Debug(err)
    },
  )
}
