package main

import (
	"log"
	"src/controllers"
	"src/core"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func resolve_command(bot telegram.BotAPI, update telegram.Update) {
	switch update.Message.Command() {
	case "help":
		controllers.Help(bot, update)
	case "roll":
		controllers.Roll_random_up_to(bot, update)
	}
}

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
		commands_map["help"] = controllers.Help
		commands_map["roll"] = controllers.Roll_random_up_to

		commands_map[update.Message.Command()](bot, update)

		resolve_command(bot, update)
	}

	log.Print("Bot is down")
}
