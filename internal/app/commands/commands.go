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
