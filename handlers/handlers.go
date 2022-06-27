package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "./web/main/index.html", gin.H{})
}
