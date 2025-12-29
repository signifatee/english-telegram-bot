package botService

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (t *TelegramService) CheckUserStatus(message *tgbotapi.Message) (bool, error) {
	user, err := t.repo.GetUser(strconv.FormatInt(message.Chat.ID, 10))
	if err != nil {
		return false, err
	}

	if user.Name != "" {
		return true, nil
	}
	return false, nil

}
