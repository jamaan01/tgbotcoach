package handler

import (
	"log/slog"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HadlerStart(bh *th.BotHandler, bot *telego.Bot, videoID string) {
	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		chatID := tu.ID(update.Message.Chat.ID)

		keyboard := tu.Keyboard(
			tu.KeyboardRow(tu.KeyboardButton("урок 1"), tu.KeyboardButton("урок 2")),
			tu.KeyboardRow(tu.KeyboardButton("урок 3"), tu.KeyboardButton("урок 4")),
			tu.KeyboardRow(tu.KeyboardButton("урок 5"), tu.KeyboardButton("урок 6")),
			tu.KeyboardRow(tu.KeyboardButton("урок 7"), tu.KeyboardButton("урок 8")),
		).WithResizeKeyboard()

		_, err := bot.SendVideoNote(ctx, tu.VideoNote(
			chatID,
			tu.FileFromID(videoID),
		))
		if err != nil {
			slog.Error("Не работает команда Start", "error", err)
		}

		message := tu.Message(
			chatID,
			"Тільки для тебе я зібрала усі ці відео в одному місці",
		).WithReplyMarkup(keyboard)
		_, err1 := bot.SendMessage(ctx, message)
		if err1 != nil {
			slog.Error("Не отправилось сообщение", "error", err1)
		}
		return nil
	}, th.CommandEqual("start"))
}
