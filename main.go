package main

import (
	"log"
	"os"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	publicGroup := router.Group("/")
	{
		publicGroup.Static("/", "./web/main")
		//publicGroup.Static("/cv", "./web/cv")
	}

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("kaze.live" + getPort()),
		Cache:      autocert.DirCache("./cert"),
	}

	log.Fatal(autotls.RunWithManager(router, &m))
	//router.RunTLS(getPort(), "./cert/server.pem", "./cert/server.key")
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "443"
	}
	return ":" + port
}

/*
func loadTls() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}*/
