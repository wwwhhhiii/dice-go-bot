package main

import (
	"diceBot/controllers"
	"diceBot/core"
	"log"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	settings := *core.NewSettings()
	bot, updates := core.Startup(settings)
	commands_map := make(map[string]func(bot telegram.BotAPI, update telegram.Update) string)

	log.Print("Bot is up")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		// map commands to corresponding controllers
		commands_map["help"] = controllers.Help
		commands_map["roll"] = controllers.Roll_random_up_to

		// Execute command
		result := commands_map[update.Message.Command()](bot, update)

		if result == "" {
			continue
		}

		msg := telegram.NewMessage(update.Message.Chat.ID, result)
		if _, err := bot.Send(msg); err != nil {
			log.Print(err)
		}
	}
}
