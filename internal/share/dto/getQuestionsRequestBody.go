package dto

type GetQuestionsRequestBody struct {
	Id string `json:"id" binding:"required"`
}
