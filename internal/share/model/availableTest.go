package model

type AvailableTest struct {
	ChatId string `json:"externalId" binding:"required" db:"chat_id"`
	TestId string `json:"testId" binding:"required" db:"test_id"`
}
