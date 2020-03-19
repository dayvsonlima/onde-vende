package config

import (
  "products/app/controllers/example_controller"
  "github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
  routes := gin.Default()

  routes.GET("/examples", example_controller.Index)

  return routes
}
