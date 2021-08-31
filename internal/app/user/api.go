package user

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "optim_22_app/pkg/log"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *gin.RouterGroup, service Service, logger log.Logger) { //, authHandler routing.Handler
  res := resource{service, logger}

//  r.Get("/user/<id>", res.get)
//  r.Get("/user", res.query)
//
//  //r.Use(authHandler)
//
  r.POST("/user", res.stubCreateUser)
//  r.Put("/user/<id>", res.update)
//  r.Delete("/user/<id>", res.delete)
}

type resource struct {
  service Service
  logger  log.Logger
}

func (r resource) stubCreateUser(c *gin.Context) { c.Status(http.StatusCreated) }
