package router

import (
	"assignment3/controllers"

	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLFiles("index.html")

	router.GET("/", controllers.SetStatus)

	return router
}
