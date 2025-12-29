package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
)

type UserStatisticRepository struct {
	db *sqlx.DB
}

func NewUserStatisticRepository(db *sqlx.DB) *UserStatisticRepository {
	return &UserStatisticRepository{db: db}
}

func (u *UserStatisticRepository) GetUserTestStatistic(chatId string, testId int) (*model.UserStatistic, error) {
	query := fmt.Sprintf("SELECT * FROM user_statistic WHERE chat_id='%v' AND test_id = %v;", chatId, testId)
	logrus.Info(query)
	var userStatistic []*model.UserStatistic
	err := u.db.Select(&userStatistic, query)
	if len(userStatistic) == 0 {
		return nil, errors.New(fmt.Sprintf("Не удалось найти статистику пользователя %v для теста %v", chatId, testId))
	}
	if err != nil {
		logrus.Errorf("Не удалось найти статистику пользователя %v для теста %v из-за ошибки: %v", chatId, testId, err)
		return nil, err
	}
	return userStatistic[0], nil
}

func (u *UserStatisticRepository) UpdateUserTestStatistic(statistic *model.UserStatistic) error {
	query := fmt.Sprintf("INSERT INTO user_statistic (chat_id, test_id, questions_number, correct_answers_number) VALUES ($1, $2, $3, $4) ON CONFLICT (chat_id, test_id) DO UPDATE SET questions_number = EXCLUDED.questions_number, correct_answers_number=EXCLUDED.correct_answers_number")
	logrus.Info(query)
	_, err := u.db.Exec(query,
		statistic.ChatId,
		statistic.TestId,
		statistic.QuestionsNumber,
		statistic.CorrectAnswersNumber,
	)
	if err != nil {
		logrus.Errorf("Ошибка при добавлении статистики юзера %v: %v", statistic.ChatId, err)
		return err
	}
	return nil
}
