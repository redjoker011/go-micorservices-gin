package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/", ShowIndexPage)

	// Create a request
	req, _ := http.NewRequest("GET", "/", nil)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

func TestShowArticlePageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/articles/:id", GetArticle)

	req, _ := http.NewRequest("GET", "/articles/1", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		// test title
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Article 1</title>") > 0

		return statusOK && pageOK
	})
}

// func TestNotFoundArticleUnauthenticated(t *testing.T) {
// 	r := getRouter(true)
// 	r.GET("/articles/:id", GetArticle)

// 	req, _ := http.NewRequest("GET", "/articles/3", nil)

// 	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
// 		return w.Code == http.StatusNotFound
// 	})
// }
