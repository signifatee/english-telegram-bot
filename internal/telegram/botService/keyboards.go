package botService

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/share/dto"
	"gitlab.com/english-vocab/telegram-bot/internal/share/model"
	"strconv"
)

var returnMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Меню"),
	),
)

var MenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Пройти тест"),
		tgbotapi.NewKeyboardButton("Статистика по тесту"),
	),
)

func CreateChooseTestKeyboard(tests []*dto.Test) tgbotapi.ReplyKeyboardMarkup {
	size := len(tests)
	logrus.Info(size)
	var buttons []tgbotapi.KeyboardButton
	var rows [][]tgbotapi.KeyboardButton
	for _, test := range tests {
		logrus.Info(test)
		buttons = append(buttons, tgbotapi.NewKeyboardButton(strconv.Itoa(test.Id)))
		size--
	}

	for i := 0; i < len(buttons); i++ {
		if i+1 >= len(buttons) {
			row := tgbotapi.NewKeyboardButtonRow(buttons[i])
			rows = append(rows, row)
			break
		}
		row := tgbotapi.NewKeyboardButtonRow(buttons[i])
		rows = append(rows, row)
	}

	rows = append(rows, tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Меню")))

	return tgbotapi.NewReplyKeyboard(rows...)
}

func (t *TelegramService) CreateNewQuestionKeyboard(questionId int) (tgbotapi.ReplyKeyboardMarkup, error) {
	options, err := t.repo.GetOptionsByQuestionId(questionId)
	if err != nil {
		logrus.Errorf("Ошибка при получении вариантов ответа по questionID: %v", err)
		return tgbotapi.ReplyKeyboardMarkup{}, err
	}

	if len(options) != 4 {
		return tgbotapi.ReplyKeyboardMarkup{}, errors.New(fmt.Sprintf("Неверное количество вариантов ответа: %v", len(options)))
	}

	questionKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(options[0].Name),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(options[1].Name),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(options[2].Name),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(options[3].Name),
		),
	)

	return questionKeyboard, nil
}

func CreateChooseLevelKeyboard(levels []*model.LanguageLevel) tgbotapi.ReplyKeyboardMarkup {
	size := len(levels)
	logrus.Info(size)
	var buttons []tgbotapi.KeyboardButton
	var rows [][]tgbotapi.KeyboardButton
	for _, level := range levels {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(level.Name))
		size--
	}

	for i := 0; i < len(buttons); i++ {
		if i+1 >= len(buttons) {
			row := tgbotapi.NewKeyboardButtonRow(buttons[i])
			rows = append(rows, row)
			break
		}
		row := tgbotapi.NewKeyboardButtonRow(buttons[i])
		rows = append(rows, row)
	}

	return tgbotapi.NewReplyKeyboard(rows...)
}

func CreateChooseGroupKeyboard(groups []*model.Group) tgbotapi.ReplyKeyboardMarkup {
	size := len(groups)
	logrus.Info(size)
	var buttons []tgbotapi.KeyboardButton
	var rows [][]tgbotapi.KeyboardButton
	for _, group := range groups {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(group.Name))
		size--
	}

	for i := 0; i < len(buttons); i++ {
		if i+1 >= len(buttons) {
			row := tgbotapi.NewKeyboardButtonRow(buttons[i])
			rows = append(rows, row)
			break
		}
		row := tgbotapi.NewKeyboardButtonRow(buttons[i])
		rows = append(rows, row)
	}

	return tgbotapi.NewReplyKeyboard(rows...)
}

func CreateChooseInstituteKeyboard(institutes []*model.Institute) tgbotapi.ReplyKeyboardMarkup {
	size := len(institutes)
	logrus.Info(size)
	var buttons []tgbotapi.KeyboardButton
	var rows [][]tgbotapi.KeyboardButton
	for _, institute := range institutes {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(institute.Name))
		size--
	}

	for i := 0; i < len(buttons); i++ {
		if i+1 >= len(buttons) {
			row := tgbotapi.NewKeyboardButtonRow(buttons[i])
			rows = append(rows, row)
			break
		}
		row := tgbotapi.NewKeyboardButtonRow(buttons[i])
		rows = append(rows, row)
	}

	return tgbotapi.NewReplyKeyboard(rows...)
}
