package controllers

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Roll_random_up_to(bot telegram.BotAPI, update telegram.Update) string {
	args := strings.Fields(update.Message.CommandArguments())
	if len(args) == 0 {
		return ""
	}

	topBorder, err := strconv.Atoi(args[0])

	if err != nil {
		return ""
	}
	if topBorder <= 0 {
		return ""
	}
	if topBorder > 10_000 {
		return ""
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	result := fmt.Sprintf("%v", random.Intn(topBorder))

	return result
}
