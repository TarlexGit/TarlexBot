package main

import (
	"fmt"
	"log"

	s "TarlexBot/settings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
	),
)

func main() {
	conf := s.New()
	// fmt.Println(conf.AdminChatId)

	admin := conf.AdminChatId
	token := conf.TGKey

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	var chats = make(map[string]int64)

	fmt.Println(chats)
	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}
		if update.Message != nil { // If we got a message
			if update.Message.Chat.ID != admin {
				// Forward all message from users to personal chat
				log.Printf("[%s | id: %d] %s", update.Message.From.UserName, update.Message.From.ID, update.Message.Text)
				chats[update.Message.From.UserName] = update.Message.Chat.ID
				chats[update.Message.From.FirstName] = update.Message.Chat.ID

				msg := tgbotapi.NewForward(admin, update.Message.Chat.ID, update.Message.MessageID)
				bot.Send(msg)
			} else {
				if update.Message.ReplyToMessage != nil {
					// ForwardSenderName
					if update.Message.ReplyToMessage.ForwardSenderName != "" {
						msg := tgbotapi.NewMessage(chats[update.Message.ReplyToMessage.ForwardSenderName], update.Message.Text)
						bot.Send(msg)
					} else {
						msg := tgbotapi.NewMessage(chats[update.Message.ReplyToMessage.ForwardFrom.UserName], update.Message.Text)
						bot.Send(msg)
					}
				} else {
					var messageTest string = "ECHO: " + update.Message.Text
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageTest)
					switch update.Message.Text {
					case "open":
						msg.ReplyMarkup = numericKeyboard
					case "close":
						msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					}

					if _, err := bot.Send(msg); err != nil {
						log.Panic(err)
					}
				}
			}
		}

	}
}
