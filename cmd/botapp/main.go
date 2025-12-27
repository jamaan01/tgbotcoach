package main

import (
	"context"
	"log/slog"
	"os"
	"tgbotcoach/internal/config"
	"tgbotcoach/internal/handler"
	"tgbotcoach/internal/telegram"

	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	bot := telegram.NewBot()
	ctx, cancel := context.WithCancel(context.Background())

	updates, err := bot.UpdatesViaLongPolling(ctx, nil)
	if err != nil {
		slog.Error("Бот не принимает айпдейты", "error", err)
		os.Exit(1)
	}

	bh, err := th.NewBotHandler(bot, updates)
	if err != nil {
		slog.Error("Ошибка инициализации хендлера", "error", err)
		os.Exit(1)
	}

	handler.HadlerStart(bh, bot, config.Load().VideoID)
	bh.Start()
	defer bh.Stop()
	defer cancel()
}
