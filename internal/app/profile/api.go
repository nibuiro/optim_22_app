package profile

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "encoding/json"
  "optim_22_app/pkg/log"
  "optim_22_app/internal/pkg/config"
  "optim_22_app/internal/pkg/utils"

)

//プロフィール操作の依存関係
type resource struct {
  config *config.Config
  service Service
  logger  log.Logger
}

//プロフィール操作についてエンドポイントを登録
func RegisterHandlers(r *gin.RouterGroup, config *config.Config, service Service, logger log.Logger) {

  rc := resource{config, service, logger}

  //取得
  r.GET("/api/profile/:userID", rc.get())
  //登録
  r.POST("/api/profile", rc.post())
  //修正
  r.PATCH("/api/profile", rc.patch())
  //削除
  r.DELETE("/api/profile", rc.delete())

}


func (rc resource) get() gin.HandlerFunc {
  return func(c *gin.Context) {
    userId := c.Param("userID")
    userProfile, err := rc.service.Get(c.Request.Context(), userId)
    if err != nil {
      rc.logger.Error(err)
      c.Status(http.StatusBadRequest)
      return 
    } else {
      if userProfileText, err := json.Marshal(userProfile); err != nil {
        rc.logger.Error(err)
        c.Status(http.StatusBadRequest)
        return
      } else {
        //c.Header("Content-Type", "application/json")
        c.JSON(http.StatusOK, userProfileText)
        return
      }
    }
  }
}


func (rc resource) post() gin.HandlerFunc {
  return func(c *gin.Context) {
    var input profile
  
    //BodyからJSONをパースして読み取る
    if err := c.BindJSON(&input); err != nil {
      rc.logger.Error(err)
      c.Status(http.StatusBadRequest)
      return 
    }
    
    //プロフィールを登録
    err := rc.service.Post(c.Request.Context(), input)
    if err != nil {
      rc.logger.Error(err)
      c.Status(http.StatusBadRequest)
      return 
    } else {
      c.Status(http.StatusCreated)
      return 
    }
  }
}


func (rc resource) patch() gin.HandlerFunc {
  return func(c *gin.Context) {
    var input profile
  
    //BodyからJSONをパースして読み取る
    if err := c.BindJSON(&input); err != nil {
      rc.logger.Error(err)
      c.Status(http.StatusBadRequest)
      return 
    }
    
    //プロフィールを登録
    err := rc.service.Post(c.Request.Context(), input)
    if err != nil {
      rc.logger.Error(err)
      c.Status(http.StatusBadRequest)
      return 
    } else {
      c.Status(http.StatusCreated)
      return 
    }
  }
}


func (rc resource) delete() gin.HandlerFunc {
  return func(c *gin.Context) {
    userId := utils.GetUserIdFromHeaderAsString(c)
    err := rc.service.Delete(c.Request.Context(), userId)
    if err != nil {
      rc.logger.Error(err)
      c.Status(http.StatusBadRequest)
      return 
    } else {
      c.Status(http.StatusOK)
      return 
    }
  }
}





