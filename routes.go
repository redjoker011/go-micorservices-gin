// Package main provides ...
package main

import "github.com/redjoker011/microservices-gin/handler"

func initializeRoutes() {
	router.GET("/", handler.ShowIndexPage)
	router.GET("/articles/:id", handler.GetArticle)
}
