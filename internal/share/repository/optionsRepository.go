package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
)

type OptionsRepository struct {
	db *sqlx.DB
}

func NewOptionsRepository(db *sqlx.DB) *OptionsRepository {
	return &OptionsRepository{db: db}
}

func (o *OptionsRepository) GetOptionById(optionId int) (*model.Option, error) {
	query := fmt.Sprintf("SELECT * FROM options WHERE option_id='%v';", optionId)
	logrus.Info(query)
	var options []*model.Option
	err := o.db.Select(&options, query)
	if len(options) == 0 {
		return nil, errors.New("нет такого варианта ответа")
	}
	if err != nil {
		logrus.Errorf("Не удалось найти вариант ответа %v из-за ошибки: %v", optionId, err)
		return nil, err
	}
	return options[0], nil
}

func (o *OptionsRepository) GetOptionsByQuestionId(questionId int) ([]*model.Option, error) {
	query := fmt.Sprintf("SELECT * FROM options WHERE options.question_id='%v';", questionId)
	logrus.Info(query)
	var options []*model.Option
	err := o.db.Select(&options, query)
	if len(options) == 0 {
		return nil, errors.New(fmt.Sprintf("нет  варианта ответа для вопроса %v", questionId))
	}
	if err != nil {
		logrus.Errorf("Не удалось найти варианты ответа для вопроса %v из-за ошибки: %v", questionId, err)
		return nil, err
	}
	return options, nil
}

func (o *OptionsRepository) PutOption(option *model.Option) error {
	query := fmt.Sprintf("INSERT INTO options (option_id, name, question_id) VALUES ($1, $2, $3) ON CONFLICT (option_id) DO UPDATE SET name = EXCLUDED.name, question_id=EXCLUDED.question_id")
	logrus.Info(query)
	_, err := o.db.Exec(query,
		option.Id,
		option.Name,
		option.QuestionId,
	)
	if err != nil {
		logrus.Errorf("Ошибка при добавлении варинта ответа %s. %s %s: %v", option.Id, option.Name, option.QuestionId, err)
		return err
	}
	return nil
}
