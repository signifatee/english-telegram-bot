package model

type Option struct {
	Id         int    `db:"option_id"`
	Name       string `db:"name"`
	QuestionId int    `db:"question_id"`
}
