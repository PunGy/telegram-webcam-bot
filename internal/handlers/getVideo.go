package handlers

import (
	"github.com/PunGy/telegram-webcam-bot/internal/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetVideoHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	config := helpers.Config()

	state := helpers.GetState()
	user := state.GetUser(update.Message.From.ID)
	user.Path = "/set-video-duration"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, config.Responses.ChooseVideoDuration)
	msg.ReplyMarkup = helpers.GetDurationKeyboard()

	bot.Send(msg)
}
