package core

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Settings struct {
	Api_key string
	Debug   string
}

func NewSettings() *Settings {
	err := godotenv.Load(("../.env"))

	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла")
	}

	settings := Settings{
		Api_key: os.Getenv("API_KEY"),
		Debug:   os.Getenv("DEBUG"),
	}

	if settings.Api_key == "" {
		log.Fatalf("В .env не предосатвлен api ключ")
	}

	return &settings
}
