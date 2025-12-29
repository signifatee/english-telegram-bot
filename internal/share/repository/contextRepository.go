package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ContextRepository struct {
	db *sqlx.DB
}

func NewContextRepository(db *sqlx.DB) *ContextRepository {
	return &ContextRepository{db: db}
}

func (c *ContextRepository) SaveContext(chatId string, context string) error {
	query := fmt.Sprintf("INSERT INTO context (context, chat_id) VALUES ($1, $2) ON CONFLICT (chat_id) DO UPDATE SET context = EXCLUDED.context")
	_, err := c.db.Exec(query,
		context,
		chatId,
	)
	if err != nil {
		logrus.Errorf("error while adding context in database: %s", err)
		return err
	}

	return nil
}

func (c *ContextRepository) GetContext(chatId string) (string, error) {
	query := fmt.Sprintf("SELECT context FROM context WHERE chat_id='%s';", chatId)
	context := []string{}
	err := c.db.Select(&context, query)
	if len(context) == 0 {
		return "", errors.New("нет контекста. Попробуйте написать /start")
	}
	return context[0], err
}
