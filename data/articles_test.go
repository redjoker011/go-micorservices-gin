// Package data provides ...
package data

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := GetAllArticles()

	if len(alist) != len(Articles) {
		t.Fail()
	}

	// Chech that each member is identical
	for i, v := range alist {
		if v.Content != Articles[i].Content ||
			v.ID != Articles[i].ID ||
			v.Title != Articles[i].Title {
			t.Fail()
			break
		}
	}
}
