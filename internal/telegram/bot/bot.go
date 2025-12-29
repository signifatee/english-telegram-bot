package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/telegram/botService"
)

const commandStart = "start"

type Bot struct {
	Bot     *tgbotapi.BotAPI
	Service *botService.Service
}

func NewBot(bot *tgbotapi.BotAPI, service *botService.Service) *Bot {
	return &Bot{
		Bot:     bot,
		Service: service,
	}
}

func (b *Bot) Start() error {

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := b.Bot.GetUpdatesChan(updateConfig)

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		context, err := b.Service.GetContext(update.Message)
		if err != nil {
			b.handleError(update.Message, err)
			continue
		}

		switch context {
		case "reg_app_name":
			msg, err := b.Service.HandleContextRegAppName(update.Message)
			if err != nil {
				logrus.Errorf("Произошла ошибка при обработке контекста reg_app_name: %v", err)
			}
			b.Bot.Send(msg)
			continue

		case "reg_app_institute":
			msg, err := b.Service.HandleContextRegAppInstitute(update.Message)
			if err != nil {
				logrus.Errorf("Произошла ошибка при обработке контекста reg_app_institute: %v", err)
			}
			b.Bot.Send(msg)
			continue

		case "reg_app_group":
			msg, err := b.Service.HandleContextRegAppGroup(update.Message)
			if err != nil {
				logrus.Errorf("Произошла ошибка при обработке контекста reg_app_group: %v", err)
			}
			b.Bot.Send(msg)
			continue

		case "reg_app_level":
			msg, err := b.Service.HandleContextRegAppLevel(update.Message)
			if err != nil {
				logrus.Errorf("Произошла ошибка при обработке контекста reg_app_level: %v", err)
			}
			b.Bot.Send(msg)
			continue

		case "reg_app_wait":
			msg, err := b.Service.HandleContextRegAppWait(update.Message)
			if err != nil {
				logrus.Errorf("Произошла ошибка при обработке контекста reg_app_wait: %v", err)
			}
			b.Bot.Send(msg)
			continue

		case "menu":
			msg, err := b.Service.HandleContextMenu(update.Message)
			if err != nil {
				logrus.Errorf("Произошла ошибка при обработке контекста menu: %v", err)
			}
			b.Bot.Send(msg)
			continue

		case "test_choose":
			msg, err := b.Service.HandleContextTestChoose(update.Message)
			if err != nil {
				logrus.Errorf("Произошла ошибка при обработке контекста test_choose: %v", err)
			}
			b.Bot.Send(msg)
			continue

		case "test_stat":
			msg, err := b.Service.HandleContextTestStat(update.Message)
			if err != nil {
				logrus.Errorf("Произошла ошибка при обработке контекста test_stat: %v", err)
			}
			b.Bot.Send(msg)
			continue

		case "get_test":
			msg, err := b.Service.HandleContextGetTest(update.Message)
			if err != nil {
				logrus.Errorf("Произошла ошибка при обработке контекста get_test: %v", err)
			}
			b.Bot.Send(msg)
			continue

		case "answering_to_test":
			msg, err := b.Service.HandleContextAnsweringToTest(update.Message)
			if err != nil {
				logrus.Errorf("Произошла ошибка при обработке контекста answering_to_test: %v", err)
			}
			b.Bot.Send(msg)
			continue
		case "choosing_stat":
			msg, err := b.Service.HandleContextChoosingStat(update.Message)
			if err != nil {
				logrus.Errorf("Произошла ошибка при обработке контекста choosing_stat: %v", err)
			}
			b.Bot.Send(msg)
			continue
		}

	}
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, "I dont know this command :(")

	switch message.Command() {
	case commandStart:
		msg, err := b.Service.HandleCommandStart(message)
		if err != nil {
			logrus.Errorf("Ошибка при обработке команды start")
			msg.Text = "Ошибка при обработке команды start"
			b.Bot.Send(msg)
		}
		b.Bot.Send(msg)
	default:
		_, err := b.Bot.Send(msg)
		return err
	}

	return nil
}

func (b *Bot) handleError(message *tgbotapi.Message, err error) {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	msg.Text = fmt.Sprintf("Что-то пошло не так при обработке вашего запроса: %s", err)

	if _, err := b.Bot.Send(msg); err != nil {
		logrus.Fatal(err)
	}
}
