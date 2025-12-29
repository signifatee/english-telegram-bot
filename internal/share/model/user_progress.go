package model

type UserProgress struct {
	ChatId     string `db:"chat_id"`
	TestId     int    `db:"test_id"`
	QuestionId int    `db:"question_id"`
	AnswerId   int    `db:"answer_id"`
}
