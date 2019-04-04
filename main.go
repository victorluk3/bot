package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var done = make(chan bool)

func main() {
	bot, err := tgbotapi.NewBotAPI("625172392:AAGTznFxi22M4m1HrAxJyRo_axd9FLmGcNk")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	go func() {
		for update := range updates {
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Text == "hi" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "hi")
				bot.Send(msg)
				continue
			}

			//really nigga?.... asi se para el bot(por el momento) y que?!..
			if update.Message.Text == "pararbot" {
				done <- true
				continue
			}

			if update.Message.Sticker != nil {
				log.Println("sticker is comming")
				msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, update.Message.Sticker.FileID)
				bot.Send(msg)
				continue
			}
			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			// msg.ReplyToMessageID = update.Message.MessageID
			if update.Message.IsCommand() {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				switch update.Message.Command() {
				case "help":
					msg.Text = "type /sayhi or /status."
				case "sayhi":
					msg.Text = "Hi :)"
				case "status":
					msg.Text = "I'm ok."
				default:
					msg.Text = "I don't know that command"
				}
				bot.Send(msg)
			}
		}
	}()

	<-done

}
