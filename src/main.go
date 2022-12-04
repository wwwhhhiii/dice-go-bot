package main

import (
	"log"
	"src/controllers"
	"src/core"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	settings := *core.NewSettings()
	bot, updates := core.Startup(settings)

	log.Print("Bot is up")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		commands_map := make(map[string]func(bot telegram.BotAPI, update telegram.Update))

		// map commands to corresponding controllers
		commands_map["help"] = controllers.Help
		commands_map["roll"] = controllers.Roll_random_up_to

		// Execute command
		commands_map[update.Message.Command()](bot, update)
	}

	log.Print("Bot is down")
}
