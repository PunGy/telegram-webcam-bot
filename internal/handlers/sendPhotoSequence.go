package handlers

import (
	"strconv"
	"time"

	"github.com/PunGy/telegram-webcam-bot/internal/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendPhotoSequence(bot *tgbotapi.BotAPI, update *tgbotapi.Update, chatId int64, size string) {
	config := helpers.Config()
	state := helpers.GetState()

	sequenceSize, err := strconv.Atoi(size)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatId, config.Responses.NumberExpected))
		return
	}
	bot.Send(tgbotapi.NewMessage(chatId, "Запланировано "+size+" фотографий"))

	for ; sequenceSize > 0; sequenceSize-- {
		img, _ := helpers.MakeWebcamImage(config.Webcam.DeviceID)

		blob := tgbotapi.FileBytes{Name: "maxim.jpg", Bytes: img}
		msg := tgbotapi.NewPhoto(chatId, blob)
		bot.Send(msg)
		if sequenceSize > 1 {
			time.Sleep(200 * time.Millisecond)
		}
	}
	state.GetUser(chatId).Path = "/"

}
