package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/dto"
	"gitlab.com/english-vocab/telegram-bot/internal/telegram/botService"
	"net/http"
	"strconv"
)

func (h *Handler) setStatusRegistrationApplication(c *gin.Context) {
	regApp := dto.RegistrationApplication{}
	err := c.ShouldBindJSON(&regApp)
	if err != nil {
		logrus.Errorf("Error while binding JSON to RegistrationApplciation: %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.SetRegistrationApplicationStatus(regApp.Status, regApp.ChatId)
	if err != nil {
		logrus.Errorf("Error while changing RegistrationApplciation status: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if regApp.Status != "Accept" {
		h.sendMessageToUsers(c)
		return
	}

	err = h.SendMessageToUserAboutApplicationStatus(regApp.ChatId, regApp.Status)
	if err != nil {
		logrus.Errorf("Ошибка при отправке сообщения пользователю: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = h.tgBot.Service.SaveContext(regApp.ChatId, "menu")
	if err != nil {
		logrus.Errorf("Ошибка при смене контекста: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Заявка на регистрацию установлена в статус: %v", regApp.Status))
}

func (h *Handler) SendMessageToUserAboutApplicationStatus(chatId string, status string) error {

	msgText := "Ошибка при изменении статуса заявки"

	switch status {
	case "Accept":
		msgText = "Ваша заявка одобрена"
	case "Reject":
		msgText = "Ваша заявка отклонена"
	case "Blocked":
		msgText = "Вы заблокированы"
	default:
		msgText = "error"
	}

	if msgText == "error" {
		return errors.New("нет такого статуса заявки")
	}

	cidInt, err := strconv.Atoi(chatId)
	if err != nil {
		logrus.Errorf("Ошибка при переводе chatid в int: %v", err)
		return err
	}
	cid := int64(cidInt)

	msg := tgbotapi.NewMessage(cid, msgText)
	msg.ReplyMarkup = botService.MenuKeyboard

	_, err = h.tgBot.Bot.Send(msg)
	if err != nil {
		logrus.Errorf("Ошибка при отправке сообщения об изменении статуса заявки пользователю: %v", err)
		return err
	}

	return nil
}
