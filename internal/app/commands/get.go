package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("not a number", args)
		return
	}
	product, err := c.productServices.Get(arg)
	if err != nil {
		log.Print("Something wrong:", product, err)
		return
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)
	c.bot.Send(msg)
}
