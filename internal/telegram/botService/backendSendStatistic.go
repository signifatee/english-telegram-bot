package botService

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/dto"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
	"net/http"
	"os"
)

func (t *TelegramService) SendStatisticsToBackend(chatId string, testId int) error {
	backendUrl := os.Getenv("BACKEND_API_URL")
	backendUrl = backendUrl + "api/service/test/create"

	userProgress, err := t.repo.GetAllAnswers(chatId, testId)
	if err != nil {
		logrus.Errorf("Ошибка при получении user progress: %v", err)
		return err
	}

	answers, err := t.UserProgressToUserAnswers(userProgress)
	if err != nil {
		logrus.Errorf("Ошибка при конвертации в userAnswer: %v", err)
		return err
	}

	jsonData, err := json.Marshal(answers)
	logrus.Info(string(jsonData))
	if err != nil {
		logrus.Errorf("Ошибка при конвертации в json при отправке статистики: %v", err)
		return err
	}

	req, err := http.NewRequest("POST", backendUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Errorf("Ошибка при создании запроса при отправке статистики: %v", err)
		return err
	}

	req.Header.Set("Authorization", os.Getenv("BACKEND_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("Ошибка при отправке запроса при отправке статистики: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.Status != "201 Created" {
		logrus.Errorf("Код ошибки при отправке статистики: %v", resp.Status)
		return errors.New("код ошибки при отправке статистики не равен 201")
	}

	return nil

}

func (t *TelegramService) UserProgressToUserAnswers(userProgress []*model.UserProgress) (*dto.UserAnswers, error) {
	length := len(userProgress)
	if length == 0 {
		return nil, errors.New("Нет отвеченных вопросов")
	}

	answers := dto.UserAnswers{
		ChatId:     userProgress[0].ChatId,
		TestId:     userProgress[0].TestId,
		AnswerList: make([]*dto.Answer, 0),
	}
	for i, up := range userProgress {
		logrus.Info(i)
		answers.AnswerList = append(answers.AnswerList, &dto.Answer{
			QuestionId: up.QuestionId,
			AnswerId:   up.AnswerId,
		})
	}
	return &answers, nil
}
