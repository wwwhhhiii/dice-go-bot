package controllers

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RollRandom(bot telegram.BotAPI, update telegram.Update) string {
	args := strings.Fields(update.Message.CommandArguments())
	if len(args) == 0 {
		return ""
	}

	topBorder, err := strconv.Atoi(args[0])
	if (err != nil) || (topBorder <= 0) || (topBorder > 10_000) {
		return ""
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	result := fmt.Sprintf("%v", random.Intn(topBorder+1))

	return result
}
