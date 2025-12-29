package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/dto"
)

type AvailableTestsForUserRepository struct {
	db *sqlx.DB
}

func NewAvailableTestsForUserRepository(db *sqlx.DB) *AvailableTestsForUserRepository {
	return &AvailableTestsForUserRepository{db: db}
}

func (a *AvailableTestsForUserRepository) GetAvailableTestsForUser(chatId string) ([]*dto.Test, error) {
	query := fmt.Sprintf("SELECT * FROM tests WHERE id IN (SELECT test_id FROM available_tests_for_user WHERE chat_id='%s');", chatId)
	logrus.Info(query)
	var tests []*dto.Test
	err := a.db.Select(&tests, query)
	if err != nil {
		return nil, err
	}
	if len(tests) == 0 {
		return nil, errors.New("Нет доступных тестов в БД")
	}
	return tests, nil
}

func (a *AvailableTestsForUserRepository) SaveAvailableTest(test *dto.Test, externalId string) error {
	query := fmt.Sprintf("INSERT INTO available_tests_for_user (chat_id, test_id) VALUES ($1, $2) ON CONFLICT (chat_id, test_id) DO UPDATE SET chat_id = EXCLUDED.chat_id, test_id = EXCLUDED.test_id")
	logrus.Info(query)
	_, err := a.db.Exec(query, externalId, test.Id)
	if err != nil {
		logrus.Errorf("Ошибка при добавлении доступного теста для пользователя в БД: %s", err)
		return err
	}

	return nil
}
