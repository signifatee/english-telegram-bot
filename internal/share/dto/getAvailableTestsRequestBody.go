package dto

type GetAvailableTestsRequestBody struct {
	ExternalId string `json:"externalId" binding:"required"`
}
