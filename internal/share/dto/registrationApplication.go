package dto

type RegistrationApplication struct {
	ChatId string `json:"externalId" binding:"required"`
	Status string `json:"status" binding:"required"`
}
