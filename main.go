package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

var gateway *PropertyGateway

func main() {
	client := resty.New()
	gateway = NewPropertyGateway("https://api.corelogic.asia", mustGetEnv("CLIENT_ID"), mustGetEnv("CLIENT_SECRET"), client)

	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("static/templates/*")

	r.GET("/", indexHandler)
	r.GET("/search", searchHandler)
	r.GET("/result/:propertyId", resultHandler)
	r.GET("/health", healthHandler)

	r.Run(":8080")
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func searchHandler(c *gin.Context) {
	suggestions, err := gateway.GetSuggestions(c.Query("q"))

	log.Println(suggestions)

	if err != nil {
		serverError(err, c)
		return
	}

	c.HTML(http.StatusOK, "suggestions.html", suggestions)
}

func resultHandler(c *gin.Context) {
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

	imagery, err := gateway.GetImagery(id)

	if err != nil {
		serverError(err, c)
		return
	}

	attributes, err := gateway.GetAttributes(id)

	if err != nil {
		serverError(err, c)
		return
	}

	model := mapValuation(valuation, imagery, attributes, c.Query("address"))

	c.HTML(http.StatusOK, "result.html", model)
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

func serverError(err error, c *gin.Context) {
	// TODO server html error page
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
