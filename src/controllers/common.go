package controllers

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Help(bot telegram.BotAPI, update telegram.Update) string {
	return "/roll N - Случайное число от 0 до N (0 < N < 10_001)"
}
