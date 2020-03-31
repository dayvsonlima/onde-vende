package config

import (
	"github.com/gin-gonic/gin"

	"website/app/controllers/home_controller"
)

func Routes() *gin.Engine {
	routes := gin.Default()
	routes.LoadHTMLGlob("app/views/**/*")

	routes.GET("/", home_controller.Index)

	return routes
}
