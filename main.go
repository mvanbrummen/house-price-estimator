package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	_ = NewPropertyGateway("https://api.corelogic.asia", os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), client)

	r := gin.Default()

	r.LoadHTMLGlob("static/templates/*")

	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	r.Run(":8080")
}
