package main

import (
	"github.com/deeplomka/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			case "list":
				listCommand(bot, update.Message, productService)
			default:
				defaultBehavior(bot, update.Message)
			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "/help - help\n"+"/list - list of products")
	_, err := bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}

func listCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message, productService *product.Service) {
	outputMsg := "Here all products list: \n"
	productsList := productService.List()

	for _, prod := range productsList {
		outputMsg += prod.Title
		outputMsg += "\n"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, outputMsg)
	_, err := bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}

func defaultBehavior(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "You wrote:"+message.Text)
	//msg.ReplyToMessageID = update.Message.MessageID

	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
