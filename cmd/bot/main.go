package main

import (
	"log"

	"github.com/beganov/test_bot/internal/app/commands"
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
	commander := commands.NewCommander(bot, productServices)
	for update := range updates {
		commander.HandlerUpdate(update)
	}
}
