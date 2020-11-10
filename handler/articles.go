package handler

import (
	"net/http"
	"strconv"

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

func GetArticle(c *gin.Context) {
	idParam := c.Param("id")
	if id, err := strconv.Atoi(idParam); err == nil {
		if article, err := data.GetArticle(id); err == nil {
			c.HTML(
				http.StatusOK,
				"article.html",
				gin.H{"title": article.Title, "payload": article},
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
