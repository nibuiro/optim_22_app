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
  "time"

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
        regexp.QuoteMeta(`INSERT INTO "comment" ("requestID","userID","date","title","body","replyID") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`),
      ).
      WillReturnRows(rows)
      suite.mock.ExpectCommit()

      t1, _ := time.Parse(time.RFC3339, "2009-11-12 21:00:57")
      insertValues := &typefile.Comment{
        RequestID: 1,
        UserID: 1,
        Date: t1,
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
      suite.mock.ExpectBegin()
      suite.mock.ExpectQuery(
        regexp.QuoteMeta(`SELECT comment.userID, user.name, comment.date, comment.title, comment.body, comment.replyID FROM "comment" INNER JOIN "user" ON comment.userID = user.ID WHERE comment.requestID = $1`),
      ).
      WillReturnRows(rows)

      var ctx context.Context
      _, err := suite.repository.Get(ctx, requestID)
    
      suite.repository.logger.Debug(err)
    },
  )
}
