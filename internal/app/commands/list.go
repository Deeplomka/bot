package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) List(message *tgbotapi.Message) {
	outputMsg := "Here all products list: \n"
	productsList := c.productService.List()

	for _, prod := range productsList {
		outputMsg += prod.Title
		outputMsg += "\n"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, outputMsg)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}
