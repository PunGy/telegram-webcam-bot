package handlers

import (
	"github.com/PunGy/telegram-webcam-bot/internal/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ShowMyselfHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	config := helpers.Config()
	img, _ := helpers.MakeWebcamImage("0")

	blob := tgbotapi.FileBytes{Name: "maxim.jpg", Bytes: img}
	msg := tgbotapi.NewPhoto(config.Basic.ClientId, blob)
	msg.Caption = config.Responses.ShowMyself

	bot.Send(msg)

	bot.Send(tgbotapi.NewMessage(config.Basic.HostId, "Done"))
}
