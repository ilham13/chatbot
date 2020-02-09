package main

import (
	"fmt"

	"github.com/ilham13/chatbot/services"
)

func main() {
	articles := services.Article{}
	// articles.getArticles()

	var article, err = articles.getArticles()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	for _, each := range article {
		fmt.Println("Title: ", each.Title)
	}

	// telegram := helpers.Telegram{}
	// telegram.Initialize()
	// telegram.ListenMessage()
}
