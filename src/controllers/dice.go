package controllers

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Roll_random_up_to(bot telegram.BotAPI, update telegram.Update) {
	args := strings.Fields(update.Message.CommandArguments())

	if len(args) == 0 {
		return
	}

	topBorder := args[0]
	topBorderInt, err := strconv.Atoi(topBorder)

	if err != nil {
		log.Panic(err)
		return
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	result := fmt.Sprintf("%v", random.Intn(topBorderInt))

	msg := telegram.NewMessage(update.Message.Chat.ID, result)

	_, err = bot.Send(msg)

	if err != nil {
		log.Panic(err)
	}
}
