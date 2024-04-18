package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Default(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "You wrote:"+message.Text)
	//msg.ReplyToMessageID = update.Message.MessageID

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
