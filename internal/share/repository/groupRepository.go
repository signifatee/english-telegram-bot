package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
)

type GroupRepository struct {
	db *sqlx.DB
}

func NewGroupRepository(db *sqlx.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (g *GroupRepository) GetAllGroups() ([]*model.Group, error) {
	query := fmt.Sprintf("SELECT * FROM groups")
	groups := []*model.Group{}
	err := g.db.Select(&groups, query)
	if len(groups) == 0 {
		return nil, errors.New("Нет групп")
	}
	return groups, err
}

func (g *GroupRepository) SaveGroup(group *model.Group) error {
	query := fmt.Sprintf("INSERT INTO groups (group_name) VALUES ($1) ON CONFLICT (group_name) DO UPDATE SET group_name = EXCLUDED.group_name")
	logrus.Info(query)
	_, err := g.db.Exec(query, group.Name)
	if err != nil {
		logrus.Errorf("Ошибка при добавлении группы в БД: %s", err)
		return err
	}

	return nil
}
