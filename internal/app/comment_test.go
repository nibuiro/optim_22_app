package comment

import (
//  "context"
  "net/http"
  "os"
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
  "optim_22_app/internal/pkg/config"
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
  
  // load application configurations
  cfg, err := config.Load("/go/src/configs/app.yaml", logger)
  if err != nil {
    logger.Errorf("failed to load application configuration: %s", err)
    os.Exit(-1)
  }

  db, mock, _ := sqlmock.New()
  suite.mock = mock
  DB, _ := gorm.Open(
    postgres.New(postgres.Config{Conn: db,}), 
    &gorm.Config{},
  )
  commentRepository := comment.StubNewRepository(DB, logger)
  suite.db = DB

  //authentication.New("localhost", "secret_key_for_refresh", "secret_key", 5, )

  commentService := comment.NewServiceStub(commentRepository, logger)

  router := gin.New()
  comment.RegisterHandlers(router.Group(""), cfg, commentService, logger)
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
        regexp.QuoteMeta(`INSERT INTO "comment" ("userID","requestID","date","title","body","replyID") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`),
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
        Body: `{"userID":1, "requestID":1, "date":"2009-11-12 21:00:57", "title":"test", "body":"test", "replyID":1}`,
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
      suite.mock.ExpectBegin()
      suite.mock.ExpectQuery(
        regexp.QuoteMeta(`SELECT comment.userID, user.name, comment.date, comment.title, comment.body, comment.replyID FROM "comment" INNER JOIN "user" ON comment.userID = user.ID WHERE comment.requestID = $1`),
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
