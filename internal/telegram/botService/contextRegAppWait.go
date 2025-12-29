package botService

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *TelegramService) HandleContextRegAppWait(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	msg.Text = fmt.Sprintf("Ваша заявка ожидает подтверждения")

	return &msg, nil
}
