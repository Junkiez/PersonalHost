package main

import (
	"main/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("web/main/*html")

	publicGroup := router.Group("/")
	{
		publicGroup.GET("/", handlers.MainPage)
		publicGroup.Static("/cv", "./web/cv")
	}

	router.Run(getPort())
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	return ":" + port
}
