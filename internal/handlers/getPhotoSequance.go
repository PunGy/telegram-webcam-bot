package handlers

import (
	"github.com/PunGy/telegram-webcam-bot/internal/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetPhotoSequenceHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	config := helpers.Config()

	state := helpers.GetState()
	user := state.GetUser(update.Message.From.ID)
	user.Path = "/get-sequence-size"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, config.Responses.ChooseSequenceSize)
	msg.ReplyMarkup = helpers.GetSequenceSizeKeyboard()

	bot.Send(msg)
}
