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
	commandsMap := make(map[string]func(bot telegram.BotAPI, update telegram.Update) string)

	log.Print("Bot is up")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Handle only commands
		if !update.Message.IsCommand() {
			continue
		}

		// map commands to corresponding controllers
		commandsMap["help"] = controllers.Help
		commandsMap["roll"] = controllers.RollRandom

		// Execute command if present in map and send message if any
		if _, isPresent := commandsMap[update.Message.Command()]; isPresent {
			result := commandsMap[update.Message.Command()](bot, update)

			if result == "" {
				continue
			}

			msg := telegram.NewMessage(update.Message.Chat.ID, result)
			if _, err := bot.Send(msg); err != nil {
				log.Print(err)
			}
		}
	}
}
