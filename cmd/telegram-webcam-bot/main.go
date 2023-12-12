package main

import (
	"github.com/PunGy/telegram-webcam-bot/internal"
	"github.com/PunGy/telegram-webcam-bot/internal/handlers"
	"github.com/PunGy/telegram-webcam-bot/internal/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func isAllowed(id int64) bool {
	cfg := helpers.Config()

	return id == cfg.Basic.HostId || id == cfg.Basic.ClientId
}

func main() {
	println("HELLO")
	cfg := helpers.Config()

	bot := internal.Initialize(cfg.Basic.BotKey, true)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if isAllowed(update.Message.From.ID) {
				internal.HandleUserMessage(bot, &update)
			} else {
				handlers.SendMessage(bot, &update, "Ухади😊", update.Message.From.ID, true)
			}
		} else if update.CallbackQuery != nil {
			if isAllowed(update.CallbackQuery.From.ID) {
				internal.HandleCallbackQuery(bot, &update)
			}
		}

	}
}
