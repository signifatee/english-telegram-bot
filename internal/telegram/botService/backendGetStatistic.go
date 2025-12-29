package botService

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/dto"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
	"io/ioutil"
	"net/http"
	"os"
)

func (t *TelegramService) GetStatisticFromBackend(getStat *dto.GetStatistic) (*model.UserStatistic, error) {
	backendUrl := os.Getenv("BACKEND_API_URL")
	backendUrl = backendUrl + "api/service/test/get-statistic"

	jsonData, err := json.Marshal(getStat)
	if err != nil {
		logrus.Errorf("Ошибка при конвертации в json при получении статистики: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", backendUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Errorf("Ошибка при создании запроса при получении статистики: %v", err)
		return nil, err
	}

	req.Header.Set("Authorization", os.Getenv("BACKEND_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("Ошибка при отправке запроса при получении статистики: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Ошибка при чтении тела ответа при получении статистики: %v", err)
		return nil, err
	}
	logrus.Infof("Статус код ответа при получении статистики: %v", resp.Status)

	statResponseBody := dto.GetStatisticResponseBody{}
	err = json.Unmarshal(body, &statResponseBody)
	if err != nil {
		logrus.Errorf("Ошибка при marshaling repsonse body при получении статистики: %v", err)
		return nil, err
	}

	questionsNumber := statResponseBody.Correct + statResponseBody.Incorrect

	userStatistic := model.UserStatistic{
		ChatId:               getStat.ChatId,
		TestId:               getStat.TestId,
		QuestionsNumber:      questionsNumber,
		CorrectAnswersNumber: statResponseBody.Correct,
	}
	return &userStatistic, nil

}
