package botService

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/dto"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
	"strconv"
)

func (t *TelegramService) HandleContextGetTest(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	if msg.Text == "Меню" {
		t.repo.SaveContext(strconv.FormatInt(message.Chat.ID, 10), "menu")
		msg.Text = "Меню"
		msg.ReplyMarkup = MenuKeyboard
		return &msg, nil
	}

	chatId := strconv.FormatInt(message.Chat.ID, 10)
	testId := msg.Text

	userHaveAccessToTest, err := t.CheckIfUserHaveAccessToTest(chatId, testId)
	if err != nil {
		msg.ReplyToMessageID = message.MessageID
		msg.Text = "Произошла ошибка при проверке доступности теста. Возможно, что вам он недоступен"
		return &msg, nil
	}

	if !userHaveAccessToTest {
		msg.ReplyToMessageID = message.MessageID
		msg.Text = "Данный тест вам недоступен"
		return &msg, nil
	}
	
	testIdInt, err := strconv.Atoi(testId)

	if err != nil {
		msg.ReplyToMessageID = message.MessageID
		msg.Text = "Произошла ошибка при конвертации testId в int"
		return &msg, nil
	}
	err = t.FillUserProgressQuestionsAndOptions(chatId, testIdInt)
	if err != nil {
		msg.ReplyToMessageID = message.MessageID
		msg.Text = "Произошла ошибка при заполнении user_progress"
		return &msg, nil
	}

	questionId, err := t.repo.GetRandomQuestionId(chatId, testIdInt)
	if err != nil {
		msg.ReplyToMessageID = message.MessageID
		msg.Text = err.Error()
		return &msg, nil
	}

	keyboard, err := t.CreateNewQuestionKeyboard(questionId)
	if err != nil {
		msg.ReplyToMessageID = message.MessageID
		msg.Text = err.Error()
		return &msg, nil
	}

	question, err := t.repo.GetQuestionById(questionId)
	if err != nil {
		msg.ReplyToMessageID = message.MessageID
		msg.Text = err.Error()
		return &msg, nil
	}

	msg.Text = fmt.Sprintf("Вопрос: \n \n%s", question.Name)
	msg.ReplyMarkup = keyboard

	err = t.repo.UpdateCurrentTestByChatId(chatId, testIdInt, questionId)
	if err != nil {
		msg.ReplyToMessageID = message.MessageID
		msg.Text = fmt.Sprintf("Произошла ошибка при обновлении currentTest: %v", err)
		return &msg, nil
	}
	t.repo.SaveContext(strconv.FormatInt(message.Chat.ID, 10), "answering_to_test")

	return &msg, nil
}

func (t *TelegramService) CheckIfUserHaveAccessToTest(chatId string, testId string) (bool, error) {
	tests, err := t.repo.GetAvailableTestsForUser(chatId)
	if err != nil {
		logrus.Errorf("ошибка при получении доступных тестов: %v", err)
		return false, err
	}

	testIdInt, _ := strconv.Atoi(testId)

	for _, test := range tests {
		if test.Id == testIdInt {
			return true, nil
		}
	}

	logrus.Errorf("Юзеру %v не доступен тест %v", chatId, testId)
	return false, nil
}

func (t *TelegramService) FillUserProgressQuestionsAndOptions(chatId string, testId int) error {
	data := dto.GetQuestionsRequestBody{Id: strconv.Itoa(testId)}
	testBody, err := t.GetQuestionsFromBackend(data)
	if err != nil {
		logrus.Errorf("Ошибка при получении вопросов из бэка: %v", err)
		return err
	}
	questions := testBody.QuestionList
	for _, question := range questions {

		err = t.repo.PutQuestion(&model.Question{QuestionId: question.QuestionId, Name: question.Name, RightOptionId: question.RightOptionId})
		if err != nil {
			logrus.Errorf("Ошибка при добавлении вопроса в БД: %v", err)
			return err
		}

		for _, option := range question.Options {
			err := t.repo.PutOption(&model.Option{
				Id:         option.Id,
				Name:       option.Name,
				QuestionId: question.QuestionId,
			})
			if err != nil {
				logrus.Errorf("Ошибка при обавлении вариантов ответа в БД: %v", err)
				return err
			}
		}

		err := t.repo.FillUserProgressWithEmptyAnswer(chatId, testId, question.QuestionId)
		if err != nil {
			logrus.Errorf("Ошибка при заполнении user progress: %v", err)
			return err
		}

	}
	return nil
}
