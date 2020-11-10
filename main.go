package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set the router as the default one provided by Gin
	router = gin.Default()
	// Load templates upfront which makes html rendering very fast
	router.LoadHTMLGlob("templates/*")

	// Initialize Routes
	initializeRoutes()

	// Start the application
	router.Run()
}
