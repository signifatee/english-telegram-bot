package botService

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (t *TelegramService) HandleContextAnsweringToTest(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	chatId := strconv.FormatInt(message.Chat.ID, 10)
	optionName := msg.Text

	currentTest, err := t.repo.GetCurrentTestByChatId(chatId)
	if err != nil {
		msg.ReplyToMessageID = message.MessageID
		msg.Text = err.Error()
		return &msg, err
	}

	//проверка, что такой вариант ответа есть
	options, err := t.repo.GetOptionsByQuestionId(currentTest.QuestionId)
	if err != nil {
		msg.ReplyToMessageID = message.MessageID
		msg.Text = err.Error()
		return &msg, err
	}

	flag := false
	var optionId int
	for _, option := range options {
		if option.Name == optionName {
			flag = true
			optionId = option.Id
		}
	}

	if !flag {
		msg.ReplyToMessageID = message.MessageID
		msg.Text = "Ответ введен некорректно"
		return &msg, err
	}

	err = t.repo.FillUserAnswer(chatId, currentTest.TestId, currentTest.QuestionId, optionId)
	if err != nil {
		msg.ReplyToMessageID = message.MessageID
		msg.Text = err.Error()
		return &msg, err
	}

	//выдача нового вопроса
	questionId, err := t.repo.GetRandomQuestionId(chatId, currentTest.TestId)
	if err != nil {
		msg.ReplyToMessageID = message.MessageID
		err := t.SendStatisticsToBackend(strconv.FormatInt(message.Chat.ID, 10), currentTest.TestId)
		if err != nil {
			msg.ReplyToMessageID = message.MessageID
			msg.Text = "Ошибка при отправке статистики в панель администратора"
			return &msg, err
		}
		t.repo.SaveContext(strconv.FormatInt(message.Chat.ID, 10), "menu")
		msg.Text = "Вы ответили на все вопросы данного теста"
		msg.ReplyMarkup = MenuKeyboard
		return &msg, err
	}

	keyboard, err := t.CreateNewQuestionKeyboard(questionId)
	if err != nil {
		msg.ReplyToMessageID = message.MessageID
		msg.Text = err.Error()
		return &msg, err
	}

	question, err := t.repo.GetQuestionById(questionId)
	if err != nil {
		msg.ReplyToMessageID = message.MessageID
		msg.Text = err.Error()
		return &msg, err
	}

	msg.Text = fmt.Sprintf("Вопрос: \n \n%s", question.Name)
	msg.ReplyMarkup = keyboard

	err = t.repo.UpdateCurrentTestByChatId(chatId, currentTest.TestId, questionId)
	if err != nil {
		msg.ReplyToMessageID = message.MessageID
		msg.Text = err.Error()
		return &msg, err
	}

	return &msg, err
}
