package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
)

type InstituteRepository struct {
	db *sqlx.DB
}

func NewInstituteRepository(db *sqlx.DB) *InstituteRepository {
	return &InstituteRepository{db: db}
}

func (i *InstituteRepository) GetInstitute(name string) (*model.Institute, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE institute_name='%s';", "institutes", name)
	logrus.Info(query)
	ins := []model.Institute{}
	err := i.db.Select(&ins, query)
	if len(ins) == 0 {
		return nil, errors.New("Нет такого института")
	}
	return &ins[0], err
}

func (i *InstituteRepository) GetAllInstitutes() (*[]model.Institute, error) {
	query := fmt.Sprintf("SELECT * FROM %s", "institutes")
	ins := []model.Institute{}
	err := i.db.Select(&ins, query)
	if len(ins) == 0 {
		return nil, errors.New("Нет институтов")
	}
	return &ins, err
}

func (i *InstituteRepository) SaveInstitute(institute *model.Institute) error {
	query := fmt.Sprintf("INSERT INTO institutes (institute_name) VALUES ($1) ON CONFLICT (institute_name) DO UPDATE SET institute_name = EXCLUDED.institute_name")
	logrus.Info(query)
	_, err := i.db.Exec(query, institute.Name)
	if err != nil {
		logrus.Errorf("Ошибка при добавлении института в БД: %s", err)
		return err
	}

	return nil
}
