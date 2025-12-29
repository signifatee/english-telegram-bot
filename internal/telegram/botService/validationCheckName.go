package botService

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func (t *TelegramService) CheckName(message *tgbotapi.Message) (bool, error) {
	name := message.Text

	words := strings.Fields(name)

	if len(words) == 3 {
		return true, nil
	} else {
		return false, errors.New("Строка не содержит 3 слова")
	}
}
