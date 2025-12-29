package dto

import "gitlab.com/english-vocab/telegram-bot/internal/share/model"

type User struct {
	ChatId        string `json:"externalId"`
	Name          string `json:"name"`
	Institute     string `json:"institute"`
	Group         string `json:"group"`
	LanguageLevel string `json:"levelEnglish"`
	TypeAccount   string `json:"typeAccount"`
}

func UserModelToDto(u *model.User) *User {
	d := User{
		ChatId:        u.ChatId,
		Name:          u.Name,
		Institute:     u.Institute,
		Group:         u.Group,
		LanguageLevel: u.LanguageLevel,
		TypeAccount:   "telegram",
	}
	return &d
}
