package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set the router as the default one provided by Gin
	router = gin.Default()
	// Load templates upfront which makes html rendering very fast
	router.LoadHTMLGlob("templates/*")

	// Define Initial Routes
	router.GET("/", func(c *gin.Context) {
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the inderx.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{"title": "Home Page"},
		)
	})

	// Start the application
	router.Run()
}
