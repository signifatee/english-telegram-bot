package dto

type Test struct {
	Id   int    `json:"id" binding:"required" db:"id"`
	Name string `json:"name" binding:"required" db:"name"`
}

type ResponseGetAvailableTestsForUser struct {
	Tests []*Test `json:"tests" binding:"required"`
}
