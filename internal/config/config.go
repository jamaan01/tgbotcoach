package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
	VideoID  string
}

func LOad() *Config {
	if err := godotenv.Load(); err != nil {
		slog.Error("Не прочитали .env", "error", err)
		os.Exit(1)
	}

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		slog.Error("Не нашелся токен телеграм бота")
		os.Exit(1)
	}

	videoID := os.Getenv("VIDEO_ID")
	if videoID == "" {
		slog.Error("Не нашелся токен кружка")
		os.Exit(1)
	}

	return &Config{
		BotToken: botToken,
		VideoID:  videoID,
	}
}
