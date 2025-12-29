package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
)

type CurrentTestRepository struct {
	db *sqlx.DB
}

func NewCurrentTestRepository(db *sqlx.DB) *CurrentTestRepository {
	return &CurrentTestRepository{db: db}
}

func (c *CurrentTestRepository) GetCurrentTestByChatId(chatId string) (*model.CurrentTest, error) {
	query := fmt.Sprintf("SELECT * FROM current_test_and_question WHERE chat_id='%v';", chatId)
	logrus.Info(query)
	var currentTests []*model.CurrentTest
	err := c.db.Select(&currentTests, query)
	if len(currentTests) == 0 {
		return nil, errors.New(fmt.Sprintf("Для пользователя %v нет текущего теста: %v", chatId, err))
	}
	if err != nil {
		logrus.Errorf("Не удалось найти текущий тест для пользователя %v из-за ошибки: %v", chatId, err)
		return nil, err
	}
	return currentTests[0], nil
}

func (c *CurrentTestRepository) UpdateCurrentTestByChatId(chatId string, testId int, questionId int) error {
	query := fmt.Sprintf("INSERT INTO current_test_and_question (chat_id, test_id, question_id) VALUES ($1, $2, $3) " +
		"ON CONFLICT (chat_id) DO UPDATE SET question_id = EXCLUDED.question_id, test_id = EXCLUDED.test_id")
	logrus.Info(query)
	_, err := c.db.Exec(query, chatId, testId, questionId)
	if err != nil {
		logrus.Errorf("Не удалось обновить текущий тест и вопрос у юзера %v на %v тест и %v вопрос. Ошибка: %v",
			chatId, testId, questionId, err)
		return err
	}
	return nil
}
