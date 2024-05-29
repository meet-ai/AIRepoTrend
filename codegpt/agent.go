package codegpt

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"jaytaylor.com/html2text"

	"github.com/tmc/langchaingo/llms/anthropic"
	gofun "github.com/x-funs/go-fun"
)

func summaryHomePage(homepage string) {
	ctx := context.Background()
	//llm, err := openai.New()
	llm, err := anthropic.New()
	if err != nil {
		log.Fatal(err)
	}
	text := getHomePageContent(homepage)
	//prompt := "What would be a good company name for a company t
	//hat makes colorful socks?"
	prompt := fmt.Sprintf("Your are an ai code expert, Please analysis the give repo homepage on github webpage, then tell your idea about this code repo, the following is <homepage> %s </homepage>", text)
	textContent := llms.TextContent{Text: prompt}
	llmMessage := llms.MessageContent{Role: llms.ChatMessageTypeHuman, Parts: []llms.ContentPart{textContent}}
	completion, err := llm.GenerateContent(ctx, []llms.MessageContent{llmMessage})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(completion.Choices[0].Content)
}

func getHomePageContent(url string) string {
	body, err := gofun.HttpGet(url)
	if err != nil {
		println(err.Error())
	}
	text, err := html2text.FromString(string(body))
	if err != nil {
		println(err.Error())
	}
	println(text)
	return text
}
