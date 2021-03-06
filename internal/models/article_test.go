package models

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := GetAllArticles()

	if len(alist) != len(ArticleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != ArticleList[i].Content ||
			v.ID != ArticleList[i].ID ||
			v.Title != ArticleList[i].Title {

			t.Fail()
			break
		}
	}

}
