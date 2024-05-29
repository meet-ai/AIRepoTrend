package codegpt

import (
	"testing"
)

func TestGetHomePage(t *testing.T) {
	getHomePageContent("https://github.com/x-funs/go-fun")
}

func TestSummary(t *testing.T) {
	summaryHomePage("https://github.com/jaytaylor/html2text")
}
