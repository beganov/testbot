package commands

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	prod := c.productServices.List()
	out := "All products: \n\n"
	for _, i := range prod {
		out = out + i.Title + " \n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, out)
	SerializedData, _ := json.Marshal(CommandData{Offset: 21})
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("14", string(SerializedData))),
	)
	c.bot.Send(msg)
}
