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
			tg.NewKeyboardButton(cfg.Keyboard.General.RequestPhotoSequence),
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

func NewNumberButton(size uint64, prefix string) tg.InlineKeyboardButton {
	str := strconv.FormatUint(size, 10)
	data := prefix + "_" + str

	return tg.InlineKeyboardButton{
		Text:         str,
		CallbackData: &data,
	}
}

func NewDurationButton(duration uint64) tg.InlineKeyboardButton {
	return NewNumberButton(duration, "duration")
}

func NewSequenceButton(size uint64) tg.InlineKeyboardButton {
	return NewNumberButton(size, "sequence")
}

func makeKeyboard(buttons []tg.InlineKeyboardButton) *tg.InlineKeyboardMarkup {
	keyboard := [][]tg.InlineKeyboardButton{
		buttons,
	}

	return &tg.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}

func GetSequenceSizeKeyboard() *tg.InlineKeyboardMarkup {
	return makeKeyboard(
		tg.NewInlineKeyboardRow(
			NewSequenceButton(3),
			NewSequenceButton(5),
			NewSequenceButton(8),
		),
	)
}

func GetDurationKeyboard() *tg.InlineKeyboardMarkup {
	return makeKeyboard(
		tg.NewInlineKeyboardRow(
			NewDurationButton(10),
			NewDurationButton(20),
			NewDurationButton(30),
			NewDurationButton(40),
		),
	)
}
