package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	_ = NewPropertyGateway("https://api.corelogic.asia", os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), client)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	r.Run(":8080")
}
