package telegram

import (
	"log/slog"
	"os"
	"tgbotcoach/internal/config"

	"github.com/mymmrac/telego"
)

func NewBot() *telego.Bot {
	bot, err := telego.NewBot(config.Load().BotToken, telego.WithDefaultDebugLogger())
	if err != nil {
		slog.Error("Не удалось создать бота", "error", err)
		os.Exit(1)
	}
	return bot
}
