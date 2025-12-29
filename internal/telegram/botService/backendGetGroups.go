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

func (t *TelegramService) GetGroupsFromBackoffice() ([]*model.Group, error) {
	backendUrl := os.Getenv("BACKEND_API_URL")
	backendUrl = backendUrl + "api/service/group/get"

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

	groups := []*model.Group{}
	err = json.Unmarshal(body, &groups)
	if err != nil {
		return nil, err
	}

	return groups, nil

}

func (t *TelegramService) SaveGroups(groups []*model.Group) error {
	var err error
	for _, group := range groups {
		err = t.repo.SaveGroup(group)
		if err != nil {
			logrus.Errorf("Ошибка при сохранении группы: %v", err)
			return err
		}
	}
	return nil
}

func (t *TelegramService) CheckGroupExists(group *model.Group, groups []*model.Group) bool {

	for _, gr := range groups {
		if gr.Name == group.Name {
			return true
		}
	}

	return false
}
