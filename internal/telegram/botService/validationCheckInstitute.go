package botService

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (t *TelegramService) CheckInstitute(message *tgbotapi.Message) (bool, error) {
	institute := message.Text
	logrus.Info(institute)
	_, err := t.repo.GetInstitute(institute)
	if err != nil {
		return false, err
	}
	return true, nil
}
