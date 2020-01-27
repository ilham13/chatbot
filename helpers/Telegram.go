package helpers

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ilham13/bot/services"
)

type Telegram struct {
	Bot             *tgbotapi.BotAPI
	State           string
	StateManagement *[]services.StateController
	LastUsername    string
	LastMessageID   int
	LastChatID      int64
}

var messageAnswer string

func (telegram *Telegram) Initialize() {
	var stateManagement []services.StateController

	telegram.StateManagement = &stateManagement
	bot, err := tgbotapi.NewBotAPI("1013961794:AAGp5iuxQtBuxCQvhXyKcdDiPxah73rVv48")
	telegram.LastUsername = ""

	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	log.Printf("ID Username %s", bot.Self.ID)

	telegram.Bot = bot
}

func (telegram *Telegram) ListenMessage() {
	var menuKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Makan siang"),
		),
	)

	var makanSiangKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Nasi Goreng", "nasi-goreng"),
			tgbotapi.NewInlineKeyboardButtonData("Ketropak", "ketropak"),
			tgbotapi.NewInlineKeyboardButtonData("Bakmie", "bakmie"),
			tgbotapi.NewInlineKeyboardButtonData("Nasi Padang", "nasi-padang"),
		),
	)

	var menuCemilanKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Cemilan"),
		),
	)
	var cemilanKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Somay", "somay"),
			tgbotapi.NewInlineKeyboardButtonData("Gorengan", "gorengan"),
			tgbotapi.NewInlineKeyboardButtonData("Bakso", "bakso"),
		),
	)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := telegram.Bot.GetUpdatesChan(u)

	for update := range updates {
		log.Println(telegram.StateManagement)
		if update.Message == nil && update.CallbackQuery == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message != nil {
			if telegram.GetState(update.Message.Chat.ID) == "end" {
				if update.CallbackQuery != nil {
					telegram.Bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, "Please callme again with /start :)"))
					continue
				}
			}
		}

		if update.Message != nil {
			if telegram.GetState(update.Message.Chat.ID) == "working" {
				if update.CallbackQuery != nil {
					telegram.Bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, "Please don't distrub me, I'm working now"))
					continue
				}
			}
		}

		if update.Message != nil {
			// messageAnswer = update.Message.Text
			messageAnswer = strings.ToLower(update.Message.Text)
			println(messageAnswer)
			switch messageAnswer {
			case "hi":
				message := "Hallo! 😉 ada yang bisa Mr. Budi bantu?"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				msg.ReplyToMessageID = update.Message.MessageID
				telegram.Bot.Send(msg)
			case "Hi":
				message := "Hallo! 😉 ada yang bisa Mr. Budi bantu?"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				msg.ReplyToMessageID = update.Message.MessageID
				telegram.Bot.Send(msg)
			case "hallo":
				message := "Hallo juga! 😉"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				msg.ReplyToMessageID = update.Message.MessageID
				telegram.Bot.Send(msg)
			case "Help":
				message := "Kamu perlu bantuan apa nih? siapa tau aja Mr. Budi bisa bantu kamu"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				telegram.Bot.Send(msg)
			case "lokasi":
				button := []tgbotapi.KeyboardButton{
					tgbotapi.NewKeyboardButtonLocation("📍 Yeay, lokasi kamu sudah ditemukan"),
				}
				replyMarkup := tgbotapi.NewReplyKeyboard(button)
				replyMarkup.OneTimeKeyboard = true

				response := tgbotapi.NewMessage(update.Message.Chat.ID, "Oke, Mr. Budi cari yaa...")
				response.BaseChat.ReplyMarkup = replyMarkup

				telegram.Bot.Send(response)
			}

			if messageAnswer == "hai" {
				message := "Hallo! 😉 ada yang bisa Mr. Budi bantu?"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				telegram.Bot.Send(msg)
			}

			if messageAnswer == "ok" || messageAnswer == "oke" {
				message := "Oke siap!"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				telegram.Bot.Send(msg)
			}

			if messageAnswer == "ada" || messageAnswer == "iya" {
				message := "Coba kamu bales 'help' deh, siapa tau ada yang bisa Mr. Budi bantu"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				telegram.Bot.Send(msg)
			}

			if messageAnswer == "help" {
				message := "Kamu perlu bantuan apa nih dari Mr. Budi? mau pesen makan? coba bales 'pengen makan' atau bales 'cemilan' klo emang lagi pengen ngemil aja"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				telegram.Bot.Send(msg)
			}

			if messageAnswer == "thanks" || messageAnswer == "thankyou" || messageAnswer == "makasih" || messageAnswer == "terima kasih" || messageAnswer == "ok thanks" {
				message := "Oke, sama-sama."
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				telegram.Bot.Send(msg)
			}

			if messageAnswer == "pagi" || messageAnswer == "morning" || messageAnswer == "Good Morning" {
				message := "Hi, Selamat pagi :)"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				telegram.Bot.Send(msg)
			}

			if messageAnswer == "sore" {
				message := "Hi, Selamat sore :)"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				telegram.Bot.Send(msg)
			}

			if messageAnswer == "sip" {
				message := "Oke!"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				telegram.Bot.Send(msg)
			}

			if messageAnswer == "kenalin dong" || messageAnswer == "kenalan" {
				message := "Masa gak kenal sama Mr. Budi :("
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				telegram.Bot.Send(msg)
			}

			if messageAnswer == "malam" || messageAnswer == "Night" || messageAnswer == "Good night" {
				message := "Hi, Selamat malam :)"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				telegram.Bot.Send(msg)
			}

			if messageAnswer == "pengen makan" || messageAnswer == "pesen makan" || messageAnswer == "laper" || messageAnswer == "makan" {
				if !telegram.IsExist(update.Message.Chat.ID) {
					telegram.CreateState(update.Message.Chat.ID)
				}
				if telegram.GetState(update.Message.Chat.ID) == "working" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please don't distrub me, I'm working now")
					telegram.Bot.Send(msg)
					continue
				} else {
					telegram.SetState(update.Message.Chat.ID, "new")
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ok, pilih dulu pesennya mau kapan ya!:")
					msg.ReplyMarkup = menuKeyboard
					msg.ReplyToMessageID = update.Message.MessageID
					telegram.Bot.Send(msg)
				}
			}

			if messageAnswer == "cemilan" || messageAnswer == "ngemil" {
				if !telegram.IsExist(update.Message.Chat.ID) {
					telegram.CreateState(update.Message.Chat.ID)
				}
				if telegram.GetState(update.Message.Chat.ID) == "working" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please don't distrub me, I'm working now")
					telegram.Bot.Send(msg)
					continue
				} else {
					telegram.SetState(update.Message.Chat.ID, "new")
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sore-sore gini emang enaknya ngemil sih")
					msg.ReplyMarkup = menuCemilanKeyboard
					msg.ReplyToMessageID = update.Message.MessageID
					telegram.Bot.Send(msg)
				}
			}

			if update.Message.IsCommand() {
				// if update.Message.Chat.ID != telegram.LastChatID {
				// 	fmt.Println("Unauthorized user : " + update.Message.From.FirstName)
				// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Who are you?!, I don't know you, so I can't serve you")
				// 	telegram.Bot.Send(msg)
				// 	continue
				// }
				switch update.Message.Command() {
				case "start":
					if !telegram.IsExist(update.Message.Chat.ID) {
						telegram.CreateState(update.Message.Chat.ID)
					}
					if telegram.GetState(update.Message.Chat.ID) == "working" {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please don't distrub me, I'm working now")
						telegram.Bot.Send(msg)
						continue
					} else {
						telegram.SetState(update.Message.Chat.ID, "new")
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ok, pilih dulu pesennya mau kapan ya!:")
						msg.ReplyMarkup = menuKeyboard
						msg.ReplyToMessageID = update.Message.MessageID
						telegram.Bot.Send(msg)
					}
				}
			} else {
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				if telegram.GetState(update.Message.Chat.ID) == "new" && update.Message.Text == "Makan siang" {
					telegram.SetState(update.Message.Chat.ID, "menu-makan")
					telegram.LastChatID = update.Message.Chat.ID
					telegram.LastUsername = update.Message.From.UserName
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Oke siap !")
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					telegram.Bot.Send(msg)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Pilih menu makanannya dulu ya..:")
					msg.ReplyMarkup = makanSiangKeyboard
					msg.ReplyToMessageID = update.Message.MessageID
					telegram.Bot.Send(msg)
					telegram.LastMessageID = update.Message.MessageID
				} else if telegram.GetState(update.Message.Chat.ID) == "new" && update.Message.Text == "Cemilan" {
					telegram.SetState(update.Message.Chat.ID, "cemilan")
					telegram.LastChatID = update.Message.Chat.ID
					telegram.LastUsername = update.Message.From.UserName
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hehe")
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					telegram.Bot.Send(msg)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Pilih menu cemilannya dulu ya..:")
					msg.ReplyMarkup = cemilanKeyboard
					msg.ReplyToMessageID = update.Message.MessageID
					telegram.Bot.Send(msg)
					telegram.LastMessageID = update.Message.MessageID
				} else {
					continue
					// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry I don't know")
					// msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					// msg.ReplyToMessageID = update.Message.MessageID
					// telegram.Bot.Send(msg)
				}
			}
		}

		if update.CallbackQuery != nil {
			query := update.CallbackQuery.Data

			if telegram.GetState(update.CallbackQuery.Message.Chat.ID) == "menu-makan" {
				switch query {
				case "ketropak":
					message := "Oke, Harga ketropaknya 15 ribu ya!, nanti saya langsung anter ke meja kerjanya ya."
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, message)
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					telegram.Bot.Send(msg)
					telegram.SetState(update.CallbackQuery.Message.Chat.ID, "end")
				case "nasi-goreng":
					message := "Oke, Harga nasi gorengnya 10 ribu ya!, nanti saya langsung anter ke meja kerjanya ya"
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, message)
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					telegram.Bot.Send(msg)
					println("nasi goreng")
					telegram.SetState(update.CallbackQuery.Message.Chat.ID, "end")
				case "bakmie":
					message := "Oke, Harga bakmienya 13 ribu ya!, nanti saya langsung anter ke meja kerjanya ya"
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, message)
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					telegram.Bot.Send(msg)
					telegram.SetState(update.CallbackQuery.Message.Chat.ID, "end")
				case "nasi-padang":
					message := "Oke, Harga nasi padangnya 20 ribu ya!, nanti saya langsung anter ke meja kerjanya ya"
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, message)
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					telegram.Bot.Send(msg)
					telegram.SetState(update.CallbackQuery.Message.Chat.ID, "end")
				}
			}

			if telegram.GetState(update.CallbackQuery.Message.Chat.ID) == "cemilan" {
				switch query {
				case "gorengan":
					message := "Oke, Harga Gorengannya 10 ribu ya!, nanti saya langsung anter ke meja kerjanya ya."
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, message)
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					telegram.Bot.Send(msg)
					telegram.SetState(update.CallbackQuery.Message.Chat.ID, "end")
				case "somay":
					message := "Oke, Harga somaynya 5 ribu ya!, nanti saya langsung anter ke meja kerjanya ya"
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, message)
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					telegram.Bot.Send(msg)
					println("nasi goreng")
					telegram.SetState(update.CallbackQuery.Message.Chat.ID, "end")
				case "bakso":
					message := "Oke, Harga bakso 13 ribu ya!, nanti saya langsung anter ke meja kerjanya ya"
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, message)
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					telegram.Bot.Send(msg)
					telegram.SetState(update.CallbackQuery.Message.Chat.ID, "end")
				}
			}
		}
	}
}

func (telegram *Telegram) CreateState(chatId int64) {
	*telegram.StateManagement = append(*telegram.StateManagement, services.StateController{ChatID: chatId, State: "new"})
}

func (telegram *Telegram) IsExist(chatId int64) bool {
	for _, n := range *telegram.StateManagement {
		if n.ChatID == chatId {
			return true
		}
	}
	return false
}

func (telegram *Telegram) GetState(chatId int64) string {
	for i, n := range *telegram.StateManagement {
		if n.ChatID == chatId {
			return (*telegram.StateManagement)[i].State
		}
	}
	return ""
}

func (telegram *Telegram) SetState(chatId int64, state string) {
	for i, n := range *telegram.StateManagement {
		if n.ChatID == chatId {
			(*telegram.StateManagement)[i].State = state
			// n.State = state
		}
	}
}
