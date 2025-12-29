package botService

import (
	"bytes"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/dto"
	"net/http"
	"os"
	"strconv"
)

func (t *TelegramService) SendUser(message *tgbotapi.Message, user *dto.User) {
	backendUrl := os.Getenv("BACKEND_API_URL")
	backendUrl = backendUrl + "api/service/registration-application/create"
	jsonData, err := json.Marshal(user)
	if err != nil {
		logrus.Errorf("cannot marshal user to json")
	}

	req, err := http.NewRequest("POST", backendUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Errorf("Ошибка при создании запроса: %v", err)
		return
	}

	req.Header.Set("Authorization", os.Getenv("BACKEND_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("Ошибка при отправке запроса: %v", err)
		return
	}
	defer resp.Body.Close()

	logrus.Infof("Статус код ответа: %v", resp.Status)

	err = t.SaveRegistrationApplicationStatus("send", strconv.FormatInt(message.Chat.ID, 10))
	if err != nil {
		logrus.Errorf("Ошибка при отправке запроса: %v", err)
	}

}
