package handlers

import (
	"os"
	"strconv"

	"github.com/PunGy/telegram-webcam-bot/internal/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetVideoFileHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update, user *tgbotapi.User, duration string) {
	config := helpers.Config()
	state := helpers.GetState()

	videoLength, err := strconv.Atoi(duration)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(user.ID, config.Responses.WrongDurationFormat))
		return
	}

	bot.Send(tgbotapi.NewMessage(user.ID, "Видео будет через "+duration+" секунд"))
	videoFilePath, _ := helpers.MakeWebcamVideo(config.Webcam.DeviceID, videoLength)

	defer os.Remove(*videoFilePath)

	msg := tgbotapi.NewVideo(user.ID, tgbotapi.FilePath(*videoFilePath))
	msg.ReplyMarkup = helpers.DefaultKeyboard(user.ID)

	state.GetUser(user.ID).Path = "/"

	bot.Send(msg)
}
