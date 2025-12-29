package model

type CurrentTest struct {
	ChatId     string `db:"chat_id"`
	TestId     int    `db:"test_id"`
	QuestionId int    `db:"question_id"`
}
