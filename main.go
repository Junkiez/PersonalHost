package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLFiles("web/main")

	publicGroup := router.Group("/")
	{
		publicGroup.GET("/", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "KAZE"})
		})
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
