package botService

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (t *TelegramService) CheckLanguageLevel(message *tgbotapi.Message) (bool, error) {
	languageLevel := message.Text
	_, err := t.repo.GetLanguageLevel(languageLevel)
	if err != nil {
		return false, err
	}
	return true, nil
}
