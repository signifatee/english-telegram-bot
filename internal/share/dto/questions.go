package dto

type Questions struct {
	QuestionId    int       `json:"questionId" binding:"required" db:"id"`
	Name          string    `json:"name" binding:"required" db:"name"`
	Options       []*Option `json:"optionList" binding:"required"`
	RightOptionId int       `json:"rightOptionId" binding:"required" db:"right_option_id"`
}
