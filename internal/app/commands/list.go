package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	prod := c.productServices.List()
	out := "All products: \n\n"
	for _, i := range prod {
		out = out + i.Title + " \n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, out)
	c.bot.Send(msg)
}
