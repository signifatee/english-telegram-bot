package botService

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.com/english-vocab/telegram-bot/internal/share/dto"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
	"strconv"
	"strings"
)

func (t *TelegramService) HandleContextRegAppLevel(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	message.Text = strings.ToUpper(message.Text)
	userMessage := message.Text

	_, err := t.CheckLanguageLevel(message)
	if err != nil {
		msg.Text = "Такого уровня нет среди доступных"
		return &msg, err
	}

	t.repo.SaveTmpUserColumn("language_level", userMessage, strconv.FormatInt(message.Chat.ID, 10))
	t.repo.SaveContext(strconv.FormatInt(message.Chat.ID, 10), "reg_app_wait")

	var user = [1]model.User{}
	user[0], err = t.repo.GetTmpUser(strconv.FormatInt(message.Chat.ID, 10))
	if err != nil {
		return nil, err
	}

	t.repo.SaveUser(&user[0])
	userModel, err := t.repo.GetUser(strconv.FormatInt(message.Chat.ID, 10))
	if err != nil {

	}
	userDto := dto.UserModelToDto(&userModel)
	t.SendUser(message, userDto)
	msg.Text = fmt.Sprintf("Ваша заявка отправлена")

	return &msg, nil
}
