package dto

type UserAnswers struct {
	ChatId     string    `json:"externalId" binding:"required" db:"chatId"`
	TestId     int       `json:"testId" binding:"required" db:"test_id"`
	AnswerList []*Answer `json:"answerList" binding:"required"`
}
