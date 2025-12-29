package botService

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.com/english-vocab/telegram-bot/internal/share/dto"
	"strconv"
)

func (t *TelegramService) HandleContextChoosingStat(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	if msg.Text == "Меню" {
		t.repo.SaveContext(strconv.FormatInt(message.Chat.ID, 10), "menu")
		msg.Text = "Меню"
		msg.ReplyMarkup = MenuKeyboard
		return &msg, nil
	}
	chatId := strconv.FormatInt(message.Chat.ID, 10)

	testId, err := strconv.Atoi(message.Text)
	if err != nil {
		msg.Text = "Ошибка при обработке id теста"
		return &msg, err
	}

	getStat := dto.GetStatistic{ChatId: chatId, TestId: testId}

	statistic, err := t.GetStatisticFromBackend(&getStat)
	if err != nil {
		return nil, err
	}

	err = t.repo.UpdateUserTestStatistic(statistic)
	if err != nil {
		msg.Text = "Ошибка при сохранении статистики"
		return &msg, err
	}
	
	testStatistic, err := t.repo.GetUserTestStatistic(chatId, testId)
	if err != nil {
		msg.Text = err.Error()
		return &msg, err
	}

	text := fmt.Sprintf("Ваша статистика по тесту %v:\n\n Всего вопросов: %v \n Правильных ответов: %v",
		testStatistic.TestId,
		testStatistic.QuestionsNumber,
		testStatistic.CorrectAnswersNumber,
	)

	msg.Text = text
	msg.ReplyMarkup = MenuKeyboard
	t.repo.SaveContext(strconv.FormatInt(message.Chat.ID, 10), "menu")

	return &msg, nil
}
