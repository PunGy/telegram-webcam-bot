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
		bot.Send(tgbotapi.NewMessage(user.ID, config.Responses.NumberExpected))
		return
	}

	bot.Send(tgbotapi.NewMessage(user.ID, "Видео будет через "+duration+" секунд"))
	videoFilePath, videoErr := helpers.MakeWebcamVideo(config.Webcam.DeviceID, videoLength)

	if videoErr != nil {
		SendMessage(bot, update, "Error occured: \""+videoErr.Error()+"\". Please try again later.", update.Message.From.ID, true)
		return
	}

	defer os.Remove(*videoFilePath)

	msg := tgbotapi.NewVideo(user.ID, tgbotapi.FilePath(*videoFilePath))
	msg.ReplyMarkup = helpers.DefaultKeyboard(user.ID)

	state.GetUser(user.ID).Path = "/"

	bot.Send(msg)

	if user.ID != config.Basic.HostId {
		// Notify the host which video was sent
		notify_msg := tgbotapi.NewVideo(config.Basic.HostId, tgbotapi.FilePath(*videoFilePath))
		notify_msg.Caption = "This video was requested from: " + user.FirstName + " " + user.LastName
		bot.Send(notify_msg)
	}
}
