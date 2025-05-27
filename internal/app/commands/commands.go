package commands

import (
	"github.com/beganov/test_bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

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
		command, ok := registeredCommands[update.Message.Command()]
		if ok {
			command(commander, update.Message)
		} else {
			commander.DefaultBehavior(update.Message)
		}
	}
}
