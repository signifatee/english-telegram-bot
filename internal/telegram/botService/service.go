package botService

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.com/english-vocab/telegram-bot/internal/share/dto"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
	"gitlab.com/english-vocab/telegram-bot/internal/share/repository"
)

type TelegramServiceInterface interface {
	HandleCommandStart(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error)

	GetContext(message *tgbotapi.Message) (string, error)
	SaveContext(chatId string, context string) error
	HandleContextRegAppName(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	HandleContextRegAppInstitute(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	HandleContextRegAppGroup(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	HandleContextRegAppWait(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	HandleContextRegAppLevel(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error)

	HandleContextMenu(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	HandleContextTestStat(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	HandleContextTestChoose(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	HandleContextGetTest(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	HandleContextAnsweringToTest(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	HandleContextChoosingStat(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error)

	CheckName(message *tgbotapi.Message) (bool, error)
	CheckLanguageLevel(message *tgbotapi.Message) (bool, error)
	CheckUserStatus(message *tgbotapi.Message) (bool, error)
	CheckInstitute(message *tgbotapi.Message) (bool, error)

	CheckIfUserHaveAccessToTest(chatId string, testId string) (bool, error)

	GetInstitutesFromBackoffice() ([]*model.Institute, error)
	GetGroupsFromBackoffice() ([]*model.Group, error)
	GetAvailableTestsForUser(data *dto.GetAvailableTestsRequestBody) ([]*dto.Test, error)
	GetQuestionsFromBackend(data dto.GetQuestionsRequestBody) (*dto.TestBody, error)

	FillUserProgressQuestionsAndOptions(chatId string, testId int) error

	SaveGroups([]*model.Group) error
	SaveInstitutes([]*model.Institute) error

	CheckGroupExists(group *model.Group, groups []*model.Group) bool

	SaveRegistrationApplicationStatus(status string, chatId string) error
	SendUser(message *tgbotapi.Message, user *dto.User)
	SendStatisticsToBackend(chatId string, testId int) error
	GetStatisticFromBackend(getStat *dto.GetStatistic) (*model.UserStatistic, error)
}

type TelegramService struct {
	repo *repository.Repository
}

func NewTelegramService(repo *repository.Repository) *TelegramService {
	return &TelegramService{repo: repo}
}

type Service struct {
	TelegramServiceInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TelegramServiceInterface: NewTelegramService(repos),
	}
}
