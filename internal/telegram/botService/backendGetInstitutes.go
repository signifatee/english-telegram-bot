package botService

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
	"io/ioutil"
	"net/http"
	"os"
)

func (t *TelegramService) GetInstitutesFromBackoffice() ([]*model.Institute, error) {
	backendUrl := os.Getenv("BACKEND_API_URL")
	backendUrl = backendUrl + "api/service/institute/get"

	req, err := http.NewRequest("POST", backendUrl, nil)
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
	logrus.Infof("Статус код ответа: %v", resp.Status)

	institutes := []*model.Institute{}
	err = json.Unmarshal(body, &institutes)
	if err != nil {
		return nil, err
	}

	return institutes, nil

}

func (t *TelegramService) SaveInstitutes(institutes []*model.Institute) error {
	var err error
	for _, institute := range institutes {
		err = t.repo.SaveInstitute(institute)
		if err != nil {
			logrus.Errorf("Ошибка при сохранении института: %v", err)
			return err
		}
	}
	return nil
}
