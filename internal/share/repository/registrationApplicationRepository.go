package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ApplicationRegistration struct {
	db *sqlx.DB
}

func NewApplicationRegistrationRepository(db *sqlx.DB) *ApplicationRegistration {
	return &ApplicationRegistration{db: db}
}

func (a *ApplicationRegistration) SetRegistrationApplicationStatus(status string, chatId string) error {
	query := fmt.Sprintf("INSERT INTO registration_application (chat_id, status) " +
		"VALUES ($1, $2) ON CONFLICT (chat_id) DO UPDATE SET status = EXCLUDED.status")
	logrus.Info(query)
	_, err := a.db.Exec(query,
		chatId,
		status,
	)
	if err != nil {
		logrus.Errorf("error while adding registration application in database: %s", err)
		return err
	}

	return nil
}
