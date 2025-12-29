package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/dto"
)

type TestsRepository struct {
	db *sqlx.DB
}

func NewTestsRepository(db *sqlx.DB) *TestsRepository {
	return &TestsRepository{db: db}
}

func (t *TestsRepository) GetTests() ([]*dto.Test, error) {
	query := fmt.Sprintf("SELECT * FROM tests;")
	logrus.Info(query)
	var tests []*dto.Test
	err := t.db.Select(&tests, query)
	if err != nil {
		return nil, err
	}
	if len(tests) == 0 {
		return nil, errors.New("в таблице нет тестов")
	}
	return tests, nil
}

func (t *TestsRepository) PutTest(test *dto.Test) error {
	query := fmt.Sprintf("INSERT INTO tests (id, name) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name")
	logrus.Info(query)
	_, err := t.db.Exec(query,
		test.Id,
		test.Name,
	)
	if err != nil {
		logrus.Errorf("Ошибка при добавлении теста %s. %s: %s", test.Id, test.Name, err)
		return err
	}
	return nil
}

func (t *TestsRepository) GetTestById(testId string) (*dto.Test, error) {
	query := fmt.Sprintf("SELECT * FROM tests WHERE id=%v;", testId)
	logrus.Info(query)
	var tests []*dto.Test
	err := t.db.Select(&tests, query)
	if err != nil {
		logrus.Errorf("Не удалось получить тест %v из-за ошибки: %v", testId, err)
		return nil, err
	}
	if len(tests) == 0 {
		return nil, errors.New("в таблице нет такого теста")
	}
	return tests[0], nil
}

func (t *TestsRepository) DeleteTestById(testId string) error {
	query := fmt.Sprintf("DELETE FROM tests WHERE id=%v", testId)
	logrus.Info(query)
	_, err := t.db.Exec(query)
	if err != nil {
		logrus.Errorf("Не удалось удалить тест %v из-за ошибки: %v", testId, err)
		return err
	}
	return nil
}
