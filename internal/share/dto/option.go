package dto

type Option struct {
	Id   int    `json:"id" binding:"required" db:"option_id"`
	Name string `json:"name" binding:"required" db:"name"`
}
