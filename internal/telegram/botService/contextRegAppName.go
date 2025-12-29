package botService

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (t *TelegramService) HandleContextRegAppName(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	userMessage := message.Text

	_, err := t.CheckName(message)
	if err != nil {
		msg.Text = "Введите ФИО в формате: Фамилия Имя Отчество"
		return &msg, err
	}

	t.repo.SaveTmpUserColumn("name", userMessage, strconv.FormatInt(message.Chat.ID, 10))
	t.repo.SaveContext(strconv.FormatInt(message.Chat.ID, 10), "reg_app_institute")

	msg.Text = fmt.Sprintf("Введите ваш институт из доступных: \n \n")

	institutes, err := t.GetInstitutesFromBackoffice()
	if err != nil {
		msg.Text = fmt.Sprintf("Не удалось получить группы из панели администратора: %v", err)
		return &msg, err
	}

	for _, institute := range institutes {
		msg.Text += fmt.Sprintf("%v \n", institute.Name)
	}

	msg.ReplyMarkup = CreateChooseInstituteKeyboard(institutes)

	return &msg, nil
}
