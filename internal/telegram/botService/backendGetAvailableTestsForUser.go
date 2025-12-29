package botService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/dto"
	"io/ioutil"
	"net/http"
	"os"
)

func (t *TelegramService) GetAvailableTestsForUser(data *dto.GetAvailableTestsRequestBody) ([]*dto.Test, error) {
	backendUrl := os.Getenv("BACKEND_API_URL")
	backendUrl = backendUrl + "api/service/test/get-available"

	jsonData, err := json.Marshal(data)
	logrus.Info(string(jsonData))
	if err != nil {
		logrus.Errorf("Ошибка при конвертации в json: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", backendUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Errorf("Ошибка при создании запроса: %v", err)
		return nil, err
	}

	req.Header.Set("Authorization", os.Getenv("BACKEND_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("Ошибка при отправке запроса: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении тела ответа:", err)
		return nil, err
	}
	logrus.Infof("Статус код ответа при получении доступных тестов: %v", resp.Status)

	logrus.Info(string(body))
	tests := dto.ResponseGetAvailableTestsForUser{}
	err = json.Unmarshal(body, &tests)
	if err != nil {
		return nil, err
	}

	return tests.Tests, nil

}
