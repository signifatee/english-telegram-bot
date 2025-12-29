package botService

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (t *TelegramService) HandleContextTestStat(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	if msg.Text == "Меню" {
		t.repo.SaveContext(strconv.FormatInt(message.Chat.ID, 10), "menu")
		msg.Text = "Меню"
		msg.ReplyMarkup = MenuKeyboard
		return &msg, nil
	}
	chatId := strconv.FormatInt(message.Chat.ID, 10)

	tests, err := t.repo.GetAvailableTestsForUser(chatId)
	if err != nil {
		msg.Text = "У вас нет доступных тестов"
		return &msg, err
	}

	msg.ReplyMarkup = CreateChooseTestKeyboard(tests)
	msg.Text = fmt.Sprintf("Выберите тест, статистику которого вы хотите посмотреть: \n\n")

	for _, test := range tests {
		msg.Text += fmt.Sprintf("%v. %s \n", test.Id, test.Name)
	}

	t.repo.SaveContext(strconv.FormatInt(message.Chat.ID, 10), "choosing_stat")
	return &msg, nil
}
