package main

import (
	"github.com/PunGy/telegram-webcam-bot/internal"
	"github.com/PunGy/telegram-webcam-bot/internal/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg := helpers.Config()

	bot := internal.Initialize(cfg.Basic.BotKey, true)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			internal.HandleUserMessage(bot, &update)
		} else if update.CallbackQuery != nil {
			internal.HandleCallbackQuery(bot, &update)
		}

	}
}
