package botService

import "github.com/sirupsen/logrus"

func (t *TelegramService) SaveRegistrationApplicationStatus(status string, chatId string) error {

	err := t.repo.SetRegistrationApplicationStatus(status, chatId)

	return err

}

func (t *TelegramService) SaveContext(chatId string, context string) error {
	err := t.repo.SaveContext(chatId, context)
	if err != nil {
		logrus.Errorf("Ошибка при сохранении конеткста в БД")
		return err
	}
	return nil
}
