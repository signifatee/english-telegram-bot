package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
)

type TmpUserRepository struct {
	db *sqlx.DB
}

func NewTmpUserRepository(db *sqlx.DB) *TmpUserRepository {
	return &TmpUserRepository{db: db}
}

func (t *TmpUserRepository) GetTmpUser(chatId string) (model.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE chat_id='%s';", "tmp_user", chatId)
	logrus.Info(query)
	user := []model.User{}
	err := t.db.Select(&user, query)
	if len(user) == 0 {
		logrus.Errorf("Не могу найти юзера в таблице tmp_user для chat_id %s", chatId)
		return model.User{}, errors.New(fmt.Sprintf("Нет такого юзера в таблице tmp_user"))
	}
	return user[0], err
}

func (t *TmpUserRepository) SaveTmpUserChatId(chatId string) error {
	query := fmt.Sprintf("INSERT INTO %s (chat_id) VALUES ($1) ON CONFLICT DO NOTHING", "tmp_user")

	logrus.Info(query)
	_, err := t.db.Exec(query, chatId)
	if err != nil {
		logrus.Errorf("error while adding tmp user's chat_id to the database: %s", err)
		return err
	}

	return nil
}

func (t *TmpUserRepository) SaveTmpUserColumn(column string, value string, chatId string) error {

	query := fmt.Sprintf("UPDATE %s SET \"%s\"='%s' WHERE chat_id='%s'", "tmp_user", column, value, chatId)

	logrus.Info(query)
	_, err := t.db.Exec(query)
	if err != nil {
		logrus.Errorf("error while adding tmp user's %s to the database: %s", column, err)
		return err
	}

	return nil
}
