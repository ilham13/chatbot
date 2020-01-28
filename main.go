package main

import (
	"github.com/ilham13/chatbot/helpers"
)

func main() {
	telegram := helpers.Telegram{}
	telegram.Initialize()
	telegram.ListenMessage()
}
