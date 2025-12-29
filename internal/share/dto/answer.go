package dto

type Answer struct {
	QuestionId int `json:"questionId" binding:"required" db:"question_id"`
	AnswerId   int `json:"answerId" binding:"required" db:"answer_id"`
}
