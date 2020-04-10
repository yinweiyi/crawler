package view

import (
	"crawler/frontend/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")
	page := model.SearchResult{
		Hits: 123,
	}
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
