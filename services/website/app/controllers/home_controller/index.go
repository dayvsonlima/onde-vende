package home_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(response *gin.Context) {
	response.HTML(http.StatusOK, "home/index.tmpl", gin.H{
		"title": "Main website",
	})
}
