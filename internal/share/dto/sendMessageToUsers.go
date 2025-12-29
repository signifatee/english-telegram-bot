package dto

type SendMessageToUsers struct {
	Message string `json:"message" binding:"required"`
}
