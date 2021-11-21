package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	gateway := NewPropertyGateway("https://api.corelogic.asia", mustGetEnv("CLIENT_ID"), mustGetEnv("CLIENT_SECRET"), client)

	r := gin.Default()

	r.LoadHTMLGlob("static/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/search", func(c *gin.Context) {
		suggestions, err := gateway.GetSuggestions(c.Query("q"))

		log.Println(suggestions)

		if err != nil {
			serverError(err, c)
			return
		}

		c.HTML(http.StatusOK, "suggestions.html", suggestions)
	})

	r.GET("/result/:propertyId", func(c *gin.Context) {
		idParam := c.Param(("propertyId"))

		id, err := strconv.Atoi(idParam)

		if err != nil {
			serverError(err, c)
			return
		}

		valuation, err := gateway.GetValuation(id)

		if err != nil {
			serverError(err, c)
			return
		}

		model := struct {
			Valuation *ValuationResponse
			Address   string
		}{
			Valuation: valuation,
			Address:   c.Query("address"),
		}

		c.HTML(http.StatusOK, "result.html", model)
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	r.Run(":8080")
}

func serverError(err error, c *gin.Context) {
	log.Println(err)
	c.JSON(500, gin.H{
		"error": err.Error(),
	})
}

func mustGetEnv(key string) string {
	e := os.Getenv(key)

	if e == "" {
		panic("Env var was not set: " + key)
	}

	return e
}
