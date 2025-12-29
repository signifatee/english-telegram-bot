package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
)

type LanguageLevelsRepository struct {
	db *sqlx.DB
}

func NewLanguageLevelsRepository(db *sqlx.DB) *LanguageLevelsRepository {
	return &LanguageLevelsRepository{db: db}
}

func (l *LanguageLevelsRepository) GetLanguageLevel(name string) (*model.LanguageLevel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE language_level_name='%s';", "language_levels", name)
	langs := []model.LanguageLevel{}
	err := l.db.Select(&langs, query)
	if len(langs) == 0 {
		return nil, errors.New("Нет такого уровня английского")
	}
	return &langs[0], err
}

func (l *LanguageLevelsRepository) GetAllLanguageLevels() ([]*model.LanguageLevel, error) {
	query := fmt.Sprintf("SELECT * FROM language_levels")
	langs := []*model.LanguageLevel{}
	err := l.db.Select(&langs, query)
	if len(langs) == 0 {
		return nil, errors.New("Нет уровней английского")
	}
	return langs, err
}
