package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var botMsg tgbotapi.MessageConfig

func UpdateHandler(u tgbotapi.UpdatesChannel, api tgbotapi.BotAPI) {

	for update := range u {

		//chatId := update.Message.Chat.ID
		msg := update.Message
		usr := msg.From
		usrName := func(usr tgbotapi.User) string {
			if usr.UserName == "" {
				return fmt.Sprintf("%s %s", usr.FirstName, usr.LastName)
			} else {
				return usr.UserName
			}
		}

		log.Printf("Message from %s", usrName(*usr))

		if msg == nil {
			continue
		}
		if msg.IsCommand() {
			commandHandler(*msg, api)
		}
	}

}
func commandHandler(msg tgbotapi.Message, api tgbotapi.BotAPI) {
	switch msg.Command() {
	case "start":
		botMsg = tgbotapi.NewMessage(msg.Chat.ID, "Start menu")
	}
	_, err := api.Send(botMsg)
	if err != nil {
		log.Println("Cant send message", err.Error())
	}
}
