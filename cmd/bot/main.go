package main

import (
	"log"

	"github.com/beganov/test_bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7626835332:AAEaA3dxfyZKaAf2Ze6ECnayNQbTgJXmRhA")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productServices := product.NewService()

	for update := range updates {

		if update.Message != nil { // If we got a message
			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			case "list":
				listCommand(bot, update.Message, productServices)
			default:
				defaultBehavior(bot, update.Message)

			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help \n /list")
	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productServices *product.Service) {
	prod := productServices.List()
	out := "All products: \n\n"
	for _, i := range prod {
		out = out + i.Title + " \n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, out)
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, inputMessage.Text)
	bot.Send(msg)
}
