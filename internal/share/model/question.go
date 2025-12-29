package model

type Question struct {
	QuestionId    int    `db:"id"`
	Name          string `db:"name"`
	RightOptionId int    `db:"right_option_id"`
}
