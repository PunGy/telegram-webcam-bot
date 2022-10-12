package internal

import (
	"strings"

	"github.com/PunGy/telegram-webcam-bot/internal/handlers"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCallbackQuery(bot *tg.BotAPI, update *tg.Update) {
	switch {
	case strings.HasPrefix(update.CallbackQuery.Data, "duration_"):
		handlers.GetVideoFileHandler(bot, update, update.CallbackQuery.From, update.CallbackQuery.Data[len("duration_"):])
	}
}
