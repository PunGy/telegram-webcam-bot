package helpers

import (
	"strconv"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DefaultKeyboard(userID int64) *tg.ReplyKeyboardMarkup {
	cfg := Config()
	keyboard := [][]tg.KeyboardButton{
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton(cfg.Keyboard.General.RequestPhoto),
			tg.NewKeyboardButton(cfg.Keyboard.General.RequestVideo),
		),
	}
	if cfg.Basic.HostId == userID {
		keyboard[0] = append(keyboard[0], tg.NewKeyboardButton(cfg.Keyboard.Host.ShowMyself))
		keyboard = append(
			keyboard,
			tg.NewKeyboardButtonRow(
				tg.NewKeyboardButton(cfg.Keyboard.Host.LittleLeave),
				tg.NewKeyboardButton(cfg.Keyboard.Host.Lunch),
				tg.NewKeyboardButton(cfg.Keyboard.Host.ImBack),
			),
			tg.NewKeyboardButtonRow(
				tg.NewKeyboardButton(cfg.Keyboard.Host.WentHome),
			),
		)
	}

	return &tg.ReplyKeyboardMarkup{
		Keyboard:       keyboard,
		ResizeKeyboard: true,
	}
}

func NewDurationButton(duration uint64) tg.InlineKeyboardButton {
	durationString := strconv.FormatUint(duration, 10)
	callbackData := "duration_" + durationString

	return tg.InlineKeyboardButton{
		Text:         durationString,
		CallbackData: &callbackData,
	}
}
func GetDurationKeyboard() *tg.InlineKeyboardMarkup {
	keyboard := [][]tg.InlineKeyboardButton{
		tg.NewInlineKeyboardRow(
			NewDurationButton(10),
			NewDurationButton(20),
			NewDurationButton(30),
			NewDurationButton(40),
		),
	}

	return &tg.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}
