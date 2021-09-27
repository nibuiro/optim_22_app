package comment

import (
//  "context"
  "net/http"
  //"os"
  "testing"
  "regexp"
  "github.com/gin-gonic/gin"
  "github.com/DATA-DOG/go-sqlmock"
  "github.com/stretchr/testify/suite"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "optim_22_app/pkg/log"
 // "optim_22_app/typefile"
  "optim_22_app/internal/app/comment"
  "optim_22_app/internal/pkg/test/v2"
)


// テストスイートの構造体
type CommentFuncIntegrationTestSuite struct {
  suite.Suite
  db *gorm.DB
  mock           sqlmock.Sqlmock
  router *gin.Engine
  logger  log.Logger
}

// テストのセットアップ
// (sqlmockをNew、Gormで発行されるクエリがモックに送られるように)
func (suite *CommentFuncIntegrationTestSuite) SetupTest() {

  logger := log.New()
  suite.logger = logger

  db, mock, _ := sqlmock.New()
  suite.mock = mock
  DB, _ := gorm.Open(
    postgres.New(postgres.Config{Conn: db,}), 
    &gorm.Config{},
  )
  commentRepository := comment.NewRepository(DB, logger)
  suite.db = DB

  //authentication.New("localhost", "secret_key_for_refresh", "secret_key", 5, )

  commentService := comment.NewService(commentRepository, logger)

  router := gin.New()
  comment.RegisterHandlers(router.Group(""), commentService, logger)
  suite.router = router
}

// テスト終了時の処理（データベース接続のクローズ）
func (suite *CommentFuncIntegrationTestSuite) TearDownTest() {
  db, _ := suite.db.DB()
  db.Close()
}

// テストスイートの実行
func TestCommentFuncIntegrationTestSuite(t *testing.T) {
  suite.Run(t, new(CommentFuncIntegrationTestSuite))
}

//コメントの投稿、取得テスト
func (suite *CommentFuncIntegrationTestSuite) TestCreate() {
  suite.Run("post a comment", func() {
      newId := 1
      rows := sqlmock.NewRows([]string{"id"}).AddRow(newId)
      suite.mock.ExpectBegin()
      suite.mock.ExpectQuery(
        regexp.QuoteMeta(`INSERT INTO "comments" ("request_id","user_id","date","title","body","reply_id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`),
      ).
      WillReturnRows(rows)
      suite.mock.ExpectCommit()
      
      //var ctx context.Context
      //err := suite.userRepository.Create(ctx, insertValues)
      testCase := test.APITestCase{
        Name: "", 
        Method: "POST", 
        URL: "/api/discussion/1", //requestID
        Header: nil, 
        Body: `{"userID":1, "requestID":1, "date":"2016-04-13T14:12:53.4242+05:30", "title":"test", "body":"test", "replyID":1}`,
        WantStatus: http.StatusCreated, 
        WantResponse: "",
      }
  
      res := test.Endpoint(suite.T(), suite.router, testCase)
    
      suite.logger.Debug(res)
    },
  )

  suite.Run("get comments", func() {
      newId := 1
      rows := sqlmock.NewRows([]string{"id"}).AddRow(newId)
     // suite.mock.ExpectBegin()
      suite.mock.ExpectQuery(
        regexp.QuoteMeta(`SELECT comments.ID, comments.RequestID, comments.UserID, user.Name, comments.Date, comments.Title, comments.Body, comments.ReplyID FROM "comments" INNER JOIN "user" ON comments.userID = user.ID WHERE comments.RequestID = $1`),
      ).
      WillReturnRows(rows)
      suite.mock.ExpectCommit()
      
      //var ctx context.Context
      //err := suite.userRepository.Create(ctx, insertValues)
      testCase := test.APITestCase{
        Name: "", 
        Method: "GET", 
        URL: "/api/discussion/1", //requestID
        Header: nil, 
        Body: "",
        WantStatus: http.StatusOK, 
        WantResponse: "",
      }
  
      res := test.Endpoint(suite.T(), suite.router, testCase)
    
      suite.logger.Debug(res)
    },
  )
}
