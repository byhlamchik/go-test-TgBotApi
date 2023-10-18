package main

import (
	handlers "TelegramBotTest/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	api, err := tgbotapi.NewBotAPI("6687687619:AAFh56a3dvSSzpc6HulC9bhkdK-eMroTkh8")
	if err != nil {
		log.Panicf("Error: %v", err.Error())
	}

	updateConfig := tgbotapi.NewUpdate(0)

	updates := api.GetUpdatesChan(updateConfig)

	handlers.UpdateHandler(updates, *api)
}
