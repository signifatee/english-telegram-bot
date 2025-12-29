package botService

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"strconv"
)

func (t *TelegramService) HandleCommandStart(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	_, err := t.repo.GetContext(strconv.FormatInt(message.Chat.ID, 10))
	if err == nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Вы уже зарегистрированы")
		return &msg, nil
	}

	err = t.repo.SaveContext(strconv.FormatInt(message.Chat.ID, 10), "reg_app_name")
	if err != nil {
		str := fmt.Sprintf("Ошибка при сохранении контекста: %v", err)
		logrus.Errorf(str)
		msg := tgbotapi.NewMessage(message.Chat.ID, str)
		return &msg, nil
	}
	context, _ := t.repo.GetContext(strconv.FormatInt(message.Chat.ID, 10))
	logrus.Infof(context)
	t.repo.SaveTmpUserChatId(strconv.FormatInt(message.Chat.ID, 10))
	msg := tgbotapi.NewMessage(message.Chat.ID, "Пожалуйста, введите ваше ФИО")

	return &msg, nil
}
