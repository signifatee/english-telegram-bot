package apiService

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
	"gitlab.com/english-vocab/telegram-bot/internal/share/repository"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetUser(chatId string) (*model.User, error) {
	user, err := u.repo.GetUser(chatId)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserService) GetAllUsers() ([]*model.User, error) {
	users, err := u.repo.GetAllUsers()
	if err != nil {
		logrus.Errorf("Ошибка при получении юзеров")
		return nil, err
	}
	return users, nil
}
