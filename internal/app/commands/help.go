package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Help(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "/help - help\n"+"/list - list of products")
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}
