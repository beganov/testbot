package commands

import (
	"github.com/beganov/test_bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot             *tgbotapi.BotAPI
	productServices *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productServices *product.Service) *Commander {
	return &Commander{
		bot:             bot,
		productServices: productServices,
	}
}

func (commander *Commander) HandlerUpdate(update tgbotapi.Update) {
	if update.Message != nil { // If we got a message

		switch update.Message.Command() {
		case "help":
			commander.Help(update.Message)
		case "get":
			commander.Get(update.Message)
		case "list":
			commander.List(update.Message)
		default:
			commander.DefaultBehavior(update.Message)

		}
	}
}
