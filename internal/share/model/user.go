package model

type User struct {
	ChatId        string `json:"externalId" db:"chat_id" binding:"required"`
	Name          string `json:"name" db:"name" binding:"required"`
	Institute     string `json:"institute" db:"institute" binding:"required"`
	Group         string `json:"group" db:"group" binding:"required"`
	LanguageLevel string `json:"levelEnglish" db:"language_level" binding:"required"`
}
