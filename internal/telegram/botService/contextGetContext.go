package botService

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"strconv"
)

func (t *TelegramService) GetContext(message *tgbotapi.Message) (string, error) {
	context, err := t.repo.GetContext(strconv.FormatInt(message.Chat.ID, 10))
	if err != nil {
		logrus.Errorf("Ошибка с получением контекста: %s", err)
	}
	return context, err
}
