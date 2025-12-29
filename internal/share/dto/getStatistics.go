package dto

type GetStatistic struct {
	ChatId string `json:"externalId" binding:"required"`
	TestId int    `json:"testId" binding:"required"`
}

type GetStatisticResponseBody struct {
	Correct   int `json:"correct" binding:"required"`
	Incorrect int `json:"inCorrect" binding:"required"`
}
