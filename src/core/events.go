package core

import telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func configure_bot(settings Settings) telegram.BotAPI {
	bot, err := telegram.NewBotAPI(settings.Api_key)
	if err != nil {
		panic(err)
	}

	if settings.Debug == "true" {
		bot.Debug = true
	}

	return *bot
}

func configure_updates(settings Settings, bot telegram.BotAPI) telegram.UpdatesChannel {
	updateConfig := telegram.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	return updates
}

func Startup(settings Settings) (telegram.BotAPI, telegram.UpdatesChannel) {
	bot := configure_bot(settings)
	updates := configure_updates(settings, bot)
	return bot, updates
}
