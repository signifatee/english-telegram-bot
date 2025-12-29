package botService

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (t *TelegramService) HandleContextMenu(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	//msg.ReplyToMessageID = message.MessageID
	msg.ReplyMarkup = returnMenu

	switch msg.Text {

	case "Пройти тест":
		t.repo.SaveContext(strconv.FormatInt(message.Chat.ID, 10), "test_choose")
		tmp, err := t.HandleContextTestChoose(message)
		if err != nil {
			msg.Text = "Произошла ошибка при выдаче тестов"
			return &msg, err
		}
		return tmp, nil

	case "Статистика по тесту":
		stat, err := t.HandleContextTestStat(message)
		if err != nil {
			msg.Text = "Произошла ошибка при выдаче тестов для просмотра статистики"
			return &msg, err
		}
		return stat, nil
	}

	msg.ReplyMarkup = MenuKeyboard
	return &msg, nil
}
