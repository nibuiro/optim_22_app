package comment

import (
  "net/http"
  "github.com/gin-gonic/gin"
)


func (rc resource) getStub() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Status(http.StatusOK)
  }
}


func (rc resource) createStub() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Status(http.StatusCreated)
  }
}







