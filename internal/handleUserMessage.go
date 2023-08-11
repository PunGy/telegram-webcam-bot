package internal

import (
	"github.com/PunGy/telegram-webcam-bot/internal/handlers"
	"github.com/PunGy/telegram-webcam-bot/internal/helpers"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUserMessage(bot *tg.BotAPI, update *tg.Update) {
	config := helpers.Config()
	state := helpers.GetState()

	user := state.GetUser(update.Message.From.ID)

	if user.Path == "/set-video-duration" {
		handlers.GetVideoFileHandler(bot, update, update.Message.From, update.Message.Text)
		return
	}
	if user.Path == "/get-sequence-size" {
		handlers.SendPhotoSequence(bot, update, update.Message.From.ID, update.Message.Text)
		return
	}

	switch update.Message.Text {
	case config.Keyboard.General.Help:
		handlers.SendMessage(bot, update, "Unknown command", update.Message.From.ID, true)
	case config.Keyboard.General.RequestPhoto:
		handlers.GetPhotoHandler(bot, update)
	case config.Keyboard.General.RequestPhotoSequence:
		handlers.GetPhotoSequenceHandler(bot, update)
	case config.Keyboard.General.RequestVideo:
		handlers.GetVideoHandler(bot, update)
	case config.Keyboard.Host.LittleLeave:
		handlers.SendMessage(bot, update, config.Responses.LittleLeave, config.Basic.ClientId, false)
	case config.Keyboard.Host.ImBack:
		handlers.SendMessage(bot, update, config.Responses.ImBack, config.Basic.ClientId, false)
	case config.Keyboard.Host.Lunch:
		handlers.SendMessage(bot, update, config.Responses.Lunch, config.Basic.ClientId, false)
	case config.Keyboard.Host.WentHome:
		handlers.SendMessage(bot, update, config.Responses.WentHome, config.Basic.ClientId, false)
	case config.Keyboard.Host.ShowMyself:
		handlers.ShowMyselfHandler(bot, update)
	default:
		handlers.SendMessage(bot, update, config.Responses.Help, update.Message.From.ID, true)
	}
}
