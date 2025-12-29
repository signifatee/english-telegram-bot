package dto

type TestBody struct {
	TestId       string       `json:"test_id" binding:"required"`
	QuestionList []*Questions `json:"questionList" binding:"required"`
}
