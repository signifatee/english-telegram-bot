package apiService

import (
	"gitlab.com/english-vocab/telegram-bot/internal/share/repository"
)

type RegistrationApplicationService struct {
	repo *repository.Repository
}

func NewRegistrationApplicationService(repo *repository.Repository) *RegistrationApplicationService {
	return &RegistrationApplicationService{repo: repo}
}

func (r *RegistrationApplicationService) SetRegistrationApplicationStatus(status string, chatId string) error {
	err := r.repo.SetRegistrationApplicationStatus(status, chatId)
	if err != nil {
		return err
	}
	return nil
}
