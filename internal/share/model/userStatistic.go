package model

type UserStatistic struct {
	ChatId               string `db:"chat_id"`
	TestId               int    `db:"test_id"`
	QuestionsNumber      int    `db:"questions_number"`
	CorrectAnswersNumber int    `db:"correct_answers_number"`
}
