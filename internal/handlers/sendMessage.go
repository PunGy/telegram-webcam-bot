package handlers

import (
	"github.com/PunGy/telegram-webcam-bot/internal/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update, message string, receiverID int64, withKeyboard bool) {
	msg := tgbotapi.NewMessage(receiverID, message)

	if withKeyboard {
		msg.ReplyMarkup = helpers.DefaultKeyboard(receiverID)
	}

	bot.Send(msg)
}
