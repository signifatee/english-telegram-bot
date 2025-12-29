package botService

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"strconv"
)

func (t *TelegramService) HandleContextRegAppInstitute(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	userMessage := message.Text

	institutes, err := t.GetInstitutesFromBackoffice()
	if err != nil {
		str := fmt.Sprintf("Ошибка при получении списка институтов: %v", err)
		logrus.Errorf(str)
		msg.Text = str
		return &msg, err
	}

	err = t.SaveInstitutes(institutes)
	if err != nil {
		str := fmt.Sprintf("Ошибка при сохранении списка институтов: %v", err)
		logrus.Errorf(str)
		msg.Text = str
		return &msg, err
	}
	_, err = t.CheckInstitute(message)
	if err != nil {
		msg.Text = "Такого института нет среди доступных"
		return &msg, err
	}

	t.repo.SaveTmpUserColumn("institute", userMessage, strconv.FormatInt(message.Chat.ID, 10))
	t.repo.SaveContext(strconv.FormatInt(message.Chat.ID, 10), "reg_app_group")

	msg.Text = fmt.Sprintf("Введите название вашей учебной группы из доступных: \n \n")

	groups, err := t.GetGroupsFromBackoffice()
	if err != nil {
		msg.Text = fmt.Sprintf("Не удалось получить группы из панели администратора: %v", err)
		return &msg, err
	}

	for _, group := range groups {
		msg.Text += fmt.Sprintf("%v \n", group.Name)
	}

	msg.ReplyMarkup = CreateChooseGroupKeyboard(groups)

	return &msg, nil
}
