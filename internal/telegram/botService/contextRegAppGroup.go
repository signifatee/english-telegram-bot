package botService

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
	"strconv"
)

func (t *TelegramService) HandleContextRegAppGroup(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	userMessage := message.Text
	group := model.Group{Name: userMessage}

	groups, err := t.GetGroupsFromBackoffice()
	if err != nil {
		str := fmt.Sprintf("Ошибка при получении групп от бэкофиса: %v", err)
		logrus.Errorf(str)
		msg.Text = str
		return &msg, err
	}
	err = t.SaveGroups(groups)
	if err != nil {
		str := fmt.Sprintf("Ошибка при сохранении групп: %v", err)
		logrus.Errorf(str)
		msg.Text = str
		return &msg, err
	}

	groupExists := t.CheckGroupExists(&group, groups)
	if !groupExists {
		msg.Text = fmt.Sprintf("Такой группы нет среди доступных")
		return &msg, nil
	}

	t.repo.SaveTmpUserColumn("group", userMessage, strconv.FormatInt(message.Chat.ID, 10))
	t.repo.SaveContext(strconv.FormatInt(message.Chat.ID, 10), "reg_app_level")

	msg.Text = fmt.Sprintf("Введите Ваш уровень английского языка")

	levels, err := t.repo.GetAllLanguageLevels()
	if err != nil {
		msg.Text = fmt.Sprintf("Не удалось получить уровни английского: %v", err)
		return &msg, err
	}

	msg.ReplyMarkup = CreateChooseLevelKeyboard(levels)

	return &msg, nil

}
