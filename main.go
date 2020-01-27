package main

import (
	"github.com/ilham13/bot/helpers"
)

func main() {
	telegram := helpers.Telegram{}
	telegram.Initialize()
	telegram.ListenMessage()
}
