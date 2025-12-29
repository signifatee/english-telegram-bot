package apiService

import (
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
	"gitlab.com/english-vocab/telegram-bot/internal/share/repository"
)

type UserServiceInterface interface {
	GetUser(chatId string) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
}

type RegistrationApplicationServiceInterface interface {
	SetRegistrationApplicationStatus(status string, chatId string) error
}

type Service struct {
	UserServiceInterface
	RegistrationApplicationServiceInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserServiceInterface:                    NewUserService(repos),
		RegistrationApplicationServiceInterface: NewRegistrationApplicationService(repos),
	}
}
