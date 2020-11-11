package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redjoker011/microservices-gin/data"
)

// Render one og HTML, JSON or CSV based on the 'Accept' header of the request
// HTML will be rendered as default
func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

func ShowIndexPage(c *gin.Context) {
	articles := data.GetAllArticles()
	// Call the HTML method of the Context to render a template
	render(
		c,
		gin.H{"title": "Home Page", "payload": articles},
		"index.html",
	)
}

func GetArticle(c *gin.Context) {
	idParam := c.Param("id")
	if id, err := strconv.Atoi(idParam); err == nil {
		if article, err := data.GetArticle(id); err == nil {
			render(
				c,
				gin.H{"title": article.Title, "payload": article},
				"article.html",
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
