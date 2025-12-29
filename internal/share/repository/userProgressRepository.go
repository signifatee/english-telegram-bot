package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
)

type UserProgressRepository struct {
	db *sqlx.DB
}

func NewUserProgressRepository(db *sqlx.DB) *UserProgressRepository {
	return &UserProgressRepository{db: db}
}

func (u *UserProgressRepository) FillUserAnswer(chatId string, testId int, questionsId, answerId int) error {
	query := fmt.Sprintf("UPDATE user_progress SET answer_id = %v WHERE chat_id = '%v' AND question_id = %v AND test_id = %v;",
		answerId,
		chatId,
		questionsId,
		testId)
	logrus.Info(query)
	_, err := u.db.Exec(query)
	if err != nil {
		logrus.Errorf("не удалось зафиксировать ответ %v пользователя %v в вопросе %v теста %v. Ошибка: %v",
			answerId,
			chatId,
			questionsId,
			testId,
			err)
		return err
	}
	return nil
}

func (u *UserProgressRepository) GetAllAnswers(chatId string, testId int) ([]*model.UserProgress, error) {
	query := fmt.Sprintf("SELECT * FROM user_progress WHERE chat_id = '%s' AND test_id = %v;", chatId, testId)
	logrus.Info(query)

	var userAnswers []*model.UserProgress
	err := u.db.Select(&userAnswers, query)
	if err != nil {
		logrus.Errorf("Ошибка при получении ответов пользователя из БД: %v", err)
		return nil, err
	}

	return userAnswers, nil
}

func (u *UserProgressRepository) GetRandomQuestionId(chatId string, testId int) (int, error) {
	query := fmt.Sprintf("SELECT question_id FROM user_progress WHERE chat_id = '%v' AND test_id = %v AND answer_id is NULL ORDER BY RANDOM();", chatId, testId)
	logrus.Info(query)
	var questionId []int
	err := u.db.Select(&questionId, query)
	if err != nil {
		logrus.Errorf("Ошибка при получении случайного вопроса для пользователя %v в тесте %v: %v", chatId, testId, err)
		return 0, err
	}

	if len(questionId) == 0 {
		logrus.Errorf("Вопросов для теста %v не осталось", testId)
		return 0, errors.New(fmt.Sprintf("Вопросов для теста %v не осталось", testId))
	}

	return questionId[0], nil
}

func (u *UserProgressRepository) FillUserProgressWithEmptyAnswer(chatId string, testId int, questionId int) error {
	query := fmt.Sprintf("INSERT INTO user_progress (chat_id, test_id, question_id) VALUES ($1, $2, $3) ON CONFLICT (chat_id, test_id, question_id) DO UPDATE SET answer_id = EXCLUDED.answer_id")
	logrus.Info(query)
	_, err := u.db.Exec(query, chatId, testId, questionId)
	if err != nil {
		logrus.Errorf("Ошибка при добавлении user progress: %s", err)
		return err
	}

	return nil
}
