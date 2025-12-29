package model

type LanguageLevel struct {
	Id   int    `db:"id"`
	Name string `db:"language_level_name"`
}
