package botService

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/dto"
	"sort"
	"strconv"
)

func (t *TelegramService) HandleContextTestChoose(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID
	if msg.Text == "Меню" {
		t.repo.SaveContext(strconv.FormatInt(message.Chat.ID, 10), "menu")
		msg.Text = "Меню"
		msg.ReplyMarkup = MenuKeyboard
		return &msg, nil
	}

	chatId := strconv.FormatInt(message.Chat.ID, 10)

	allTests, err := t.GetAllTests()
	if err != nil {
		str := fmt.Sprintf("Ошибка при получении всех тестов из бэкофиса: %v", err)
		logrus.Errorf(str)
		msg.Text = str
		return &msg, err
	}

	for _, test := range allTests {
		err := t.repo.PutTest(test)
		if err != nil {
			str := fmt.Sprintf("Ошибка при получении сохранении всех тестов из бэкофиса: %v", err)
			logrus.Errorf(str)
			msg.Text = str
			return &msg, err
		}
	}

	data := &dto.GetAvailableTestsRequestBody{ExternalId: chatId}

	tests, err := t.GetAvailableTestsForUser(data)
	if err != nil {
		str := fmt.Sprintf("Ошибка при получении доступных тестов из бэкофиса: %v", err)
		logrus.Errorf(str)
		msg.Text = str
		return &msg, err
	}

	for _, test := range tests {
		logrus.Info(&test)
		err := t.repo.SaveAvailableTest(test, chatId)
		if err != nil {
			str := fmt.Sprintf("Ошибка при сохранении доступных тестов: %v", err)
			logrus.Errorf(str)
			msg.Text = str
			return &msg, err
		}
	}

	//tests, err = t.repo.GetAvailableTestsForUser(strconv.FormatInt(message.Chat.ID, 10))
	//if err != nil {
	//	msg.Text = "У вас нет доступных тестов"
	//	return &msg, err
	//}

	msg.Text = "Выберите один из доступных вам тестов: \n"

	sort.Slice(tests, func(i, j int) bool {
		return tests[i].Id < tests[j].Id
	})

	for _, test := range tests {
		msg.Text += fmt.Sprintf("%v. %s \n", test.Id, test.Name)
	}

	msg.ReplyMarkup = CreateChooseTestKeyboard(tests)

	t.repo.SaveContext(strconv.FormatInt(message.Chat.ID, 10), "get_test")
	return &msg, nil
}
