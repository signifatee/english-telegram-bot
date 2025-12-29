package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
)

type QuestionsRepository struct {
	db *sqlx.DB
}

func NewQuestionsRepository(db *sqlx.DB) *QuestionsRepository {
	return &QuestionsRepository{db: db}
}

func (q *QuestionsRepository) GetQuestionById(questionId int) (*model.Question, error) {
	query := fmt.Sprintf("SELECT * FROM questions WHERE id='%v';", questionId)
	logrus.Info(query)
	var questions []*model.Question
	err := q.db.Select(&questions, query)
	if len(questions) == 0 {
		return nil, errors.New("Нет такого вопроса")
	}
	if err != nil {
		logrus.Errorf("Не удалось найти вопрос %v из-за ошибки: %v", questionId, err)
		return nil, err
	}
	return questions[0], err
}

func (q *QuestionsRepository) PutQuestion(question *model.Question) error {
	query := fmt.Sprintf("INSERT INTO questions (id, name, right_option_id) VALUES ($1, $2, $3) ON CONFLICT (id) " +
		"DO UPDATE SET name = EXCLUDED.name, right_option_id = EXCLUDED.right_option_id;")
	logrus.Info(query)
	_, err := q.db.Exec(query,
		question.QuestionId,
		question.Name,
		question.RightOptionId,
	)
	if err != nil {
		logrus.Errorf("Ошибка при добавлении вопроса %s. %s с правильным ответом %v: %s",
			question.QuestionId,
			question.Name,
			question.RightOptionId,
			err,
		)
		return err
	}
	return nil
}
