package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redjoker011/microservices-gin/data"
)

func ShowIndexPage(c *gin.Context) {
	articles := data.GetAllArticles()
	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the inderx.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{"title": "Home Page", "payload": articles},
	)
}
