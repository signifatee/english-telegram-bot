package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/dto"
	"net/http"
	"strconv"
)

func (h *Handler) sendMessageToUsers(c *gin.Context) {
	var sendInfo dto.SendMessageToUsers
	err := c.ShouldBindJSON(&sendInfo)
	if err != nil {
		logrus.Errorf("Error while binding JSON to SendMessageToUsers: %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	users, err := h.services.GetAllUsers()
	if err != nil {
		logrus.Errorf("Ошибка при получении списка юзеров: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	strError := ""

	for _, user := range users {
		err := h.SendMessageToUserFromBackend(user.ChatId, sendInfo.Message)
		if err != nil {
			str := fmt.Sprintf("Error while sending message to %s: %s", user.ChatId, err.Error())
			logrus.Errorf(str)
			strError += str + "\n"
		}
	}

	c.JSON(http.StatusOK, strError)
	return
}

func (h *Handler) SendMessageToUserFromBackend(chatId string, message string) error {

	msgText := message

	cidInt, err := strconv.Atoi(chatId)
	if err != nil {
		logrus.Errorf("Ошибка при переводе chatid в int: %v", err)
		return err
	}
	cid := int64(cidInt)

	logrus.Infof("Отправка сообщения %s юзеру %s", chatId, message)
	msg := tgbotapi.NewMessage(cid, msgText)
	//msg.ReplyMarkup = botService.MenuKeyboard

	_, err = h.tgBot.Bot.Send(msg)
	if err != nil {
		logrus.Errorf("Ошибка при отправке сообщения об изменении статуса заявки пользователю: %v", err)
		return err
	}

	return nil
}
