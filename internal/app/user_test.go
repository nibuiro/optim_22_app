package user

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
  "optim_22_app/internal/app/user"
  "optim_22_app/internal/pkg/config"
  "optim_22_app/internal/pkg/test/v2"
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
  db *gorm.DB
  mock           sqlmock.Sqlmock
  router *gin.Engine
  logger  log.Logger
}

// テストのセットアップ
// (sqlmockをNew、Gormで発行されるクエリがモックに送られるように)
func (suite *UserRepositoryTestSuite) SetupTest() {

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
  userRepository := user.NewRepository(DB, logger)
  suite.db = DB

  //authentication.New("localhost", "secret_key_for_refresh", "secret_key", 5, )

  userService := user.NewService(userRepository, logger)

  router := gin.New()
  user.RegisterHandlers(router.Group(""), cfg, userService, logger)
  suite.router = router
}

// テスト終了時の処理（データベース接続のクローズ）
func (suite *UserRepositoryTestSuite) TearDownTest() {
  db, _ := suite.db.DB()
  db.Close()
}

// テストスイートの実行
func TestUserRepositoryTestSuite(t *testing.T) {
  suite.Run(t, new(UserRepositoryTestSuite))
}

// CreateのテストTestCreate
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
      
      //var ctx context.Context
      //err := suite.userRepository.Create(ctx, insertValues)
      testCase := test.APITestCase{
        Name: "register success", 
        Method: "POST", 
        URL: "/api/user", 
        Header: nil, 
        Body: `{"name":"test", "email":"test@test.test", "password":"7f83b1657ff1fc53b92dc18148a1d6fffffd4b1fa3d677284addd200126d9069"}`,
        WantStatus: http.StatusCreated, 
        WantResponse: "",
      }
  
      //ヘッダとCookie以外について検証
      res := test.Endpoint(suite.T(), suite.router, testCase)
    
      suite.logger.Debug(res)
    //  suite.logger.Debug(err)
    },
  )
}




type FakeAuthorizationService struct {
  repository interface{}
}

func (as *FakeAuthorizationService) New(args ...interface{}) (string, string) { 
  return "", ""
}

func (as *FakeAuthorizationService) Refresh(refreshToken string) string {
  return ""
}