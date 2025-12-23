package handler

import (
	"log/slog"
	"tgbotcoach/internal/config"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HadlerStart(bh *th.BotHandler, bot *telego.Bot) {
	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		chatID := tu.ID(update.Message.Chat.ID)
		_, err := bot.SendVideoNote(ctx, tu.VideoNote(
			chatID,
			tu.FileFromID(config.Load().VideoID),
		))
		if err != nil {
			slog.Error("Не работает команда Start", "error", err)
		}

		message := tu.Message(
			chatID,
			"Тільки для тебе я зібрала усі ці відео в одному місці",
		)
		_, err1 := bot.SendMessage(ctx, message)
		if err1 != nil {
			slog.Error("Не отправилось сообщение", "error", err1)
		}
		return nil
	}, th.CommandEqual("start"))
}
