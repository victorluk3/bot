package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, r.URL.Path[1:])
		})

		http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hi")
		})

		//fmt.Printf("server runing port: %v\n", os.Getenv("PORT"))
		http.ListenAndServe(os.Getenv("PORT"), nil)

	}()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TKEY"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	updates, err := bot.GetUpdatesChan(u)

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
}
