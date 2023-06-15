package handlers

import (
	"github.com/PunGy/telegram-webcam-bot/internal/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetPhotoHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	config := helpers.Config()
	img, _ := helpers.MakeWebcamImage(config.Webcam.DeviceID)

	blob := tgbotapi.FileBytes{Name: "maxim.jpg", Bytes: img}
	msg := tgbotapi.NewPhoto(update.Message.Chat.ID, blob)
	msg.ReplyMarkup = helpers.DefaultKeyboard(update.Message.From.ID)

	bot.Send(msg)

	if update.Message.Chat.ID != config.Basic.HostId {
		// Notify the host which photo was sent
		notify_msg := tgbotapi.NewPhoto(config.Basic.HostId, blob)
		notify_msg.Caption = "This photo was requested from: " + update.Message.From.FirstName + " " + update.Message.From.LastName
		bot.Send(notify_msg)
	}
}
