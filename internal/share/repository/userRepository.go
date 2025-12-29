package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) DeleteUser(chatId string) error {
	return nil
}

func (u *UserRepository) GetUser(chatId string) (model.User, error) {
	query := fmt.Sprintf("SELECT * FROM \"user\" WHERE chat_id='%s';", chatId)
	logrus.Info(query)
	user := []model.User{}
	err := u.db.Select(&user, query)
	if len(user) == 0 {
		logrus.Errorf("Не могу найти юзера в таблице user для chat_id %s", chatId)
		return model.User{}, errors.New(fmt.Sprintf("Нет такого юзера в таблице user"))
	}
	return user[0], err
}

func (u *UserRepository) SaveUser(user *model.User) error {
	query := fmt.Sprintf("INSERT INTO \"%s\" (chat_id, "+
		"name, "+
		"institute, "+
		"\"group\", "+
		"language_level) "+
		"VALUES ($1, $2, $3, $4, $5) ON CONFLICT (chat_id) "+
		"DO UPDATE SET "+
		"name = EXCLUDED.name, "+
		"institute = EXCLUDED.institute, "+
		"\"group\" = EXCLUDED.\"group\", "+
		"language_level = EXCLUDED.language_level", "user")

	logrus.Info(query)
	_, err := u.db.Exec(query,
		user.ChatId,
		user.Name,
		user.Institute,
		user.Group,
		user.LanguageLevel,
	)

	if err != nil {
		logrus.Errorf("error while adding user to the database: %s", err)
		return err
	}

	return nil

}

func (u *UserRepository) GetAllUsers() ([]*model.User, error) {
	query := fmt.Sprintf("SELECT * FROM \"user\";")
	logrus.Info(query)
	users := []*model.User{}
	err := u.db.Select(&users, query)
	if len(users) == 0 {
		logrus.Errorf("Нет юзеров")
		return nil, errors.New(fmt.Sprintf("Нет юзеров"))
	}
	return users, err
}
