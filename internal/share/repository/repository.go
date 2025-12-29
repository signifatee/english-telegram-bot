package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/english-vocab/telegram-bot/internal/share/dto"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
)

type LanguageLevelsRepositoryInterface interface {
	GetLanguageLevel(name string) (*model.LanguageLevel, error)
	GetAllLanguageLevels() ([]*model.LanguageLevel, error)
}

type ApplicationRegistrationInterface interface {
	SetRegistrationApplicationStatus(status string, chatId string) error
}

type InstituteRepositoryInterface interface {
	GetInstitute(name string) (*model.Institute, error)
	GetAllInstitutes() (*[]model.Institute, error)
	SaveInstitute(*model.Institute) error
}

type GroupRepositoryInterface interface {
	SaveGroup(group *model.Group) error
	GetAllGroups() ([]*model.Group, error)
}

type UserRepositoryInterface interface {
	SaveUser(u *model.User) error
	GetUser(chatId string) (model.User, error)
	GetAllUsers() ([]*model.User, error)
	DeleteUser(chatId string) error
}

type TmpUserRepositoryInterface interface {
	GetTmpUser(chatId string) (model.User, error)
	SaveTmpUserChatId(chatId string) error
	SaveTmpUserColumn(column string, value string, chat_id string) error
}

type ContextRepositoryInterface interface {
	SaveContext(chatId string, context string) error
	GetContext(chatId string) (string, error)
}

type AvailableTestsForUserInterface interface {
	GetAvailableTestsForUser(chatId string) ([]*dto.Test, error)
	SaveAvailableTest(test *dto.Test, externalId string) error
}

type UserProgressInterface interface {
	GetRandomQuestionId(chatId string, testId int) (int, error)
	FillUserAnswer(chatId string, testId int, questionsId, answerId int) error
	GetAllAnswers(chatId string, testId int) ([]*model.UserProgress, error)
	FillUserProgressWithEmptyAnswer(chatId string, testId int, questionId int) error
}

type TestsInterface interface {
	GetTests() ([]*dto.Test, error)
	PutTest(test *dto.Test) error
	GetTestById(testId string) (*dto.Test, error)
	DeleteTestById(testId string) error
}

type QuestionsInterface interface {
	GetQuestionById(questionId int) (*model.Question, error)
	PutQuestion(question *model.Question) error
}

type OptionsInterface interface {
	GetOptionById(optionId int) (*model.Option, error)
	GetOptionsByQuestionId(questionId int) ([]*model.Option, error)
	PutOption(option *model.Option) error
}

type CurrentTestInterface interface {
	GetCurrentTestByChatId(chatId string) (*model.CurrentTest, error)
	UpdateCurrentTestByChatId(chatId string, testId int, questionId int) error
}

type UserStatisticInterface interface {
	UpdateUserTestStatistic(statistic *model.UserStatistic) error
	GetUserTestStatistic(chatId string, testId int) (*model.UserStatistic, error)
}

type Repository struct {
	UserRepositoryInterface
	TmpUserRepositoryInterface
	ContextRepositoryInterface
	InstituteRepositoryInterface
	LanguageLevelsRepositoryInterface
	ApplicationRegistrationInterface
	AvailableTestsForUserInterface
	UserProgressInterface
	TestsInterface
	QuestionsInterface
	OptionsInterface
	CurrentTestInterface
	UserStatisticInterface
	GroupRepositoryInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepositoryInterface:           NewUserRepository(db),
		TmpUserRepositoryInterface:        NewTmpUserRepository(db),
		ContextRepositoryInterface:        NewContextRepository(db),
		InstituteRepositoryInterface:      NewInstituteRepository(db),
		LanguageLevelsRepositoryInterface: NewLanguageLevelsRepository(db),
		ApplicationRegistrationInterface:  NewApplicationRegistrationRepository(db),
		AvailableTestsForUserInterface:    NewAvailableTestsForUserRepository(db),
		UserProgressInterface:             NewUserProgressRepository(db),
		TestsInterface:                    NewTestsRepository(db),
		QuestionsInterface:                NewQuestionsRepository(db),
		OptionsInterface:                  NewOptionsRepository(db),
		CurrentTestInterface:              NewCurrentTestRepository(db),
		UserStatisticInterface:            NewUserStatisticRepository(db),
		GroupRepositoryInterface:          NewGroupRepository(db),
	}
}
