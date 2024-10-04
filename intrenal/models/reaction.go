package models

type Reaction struct {
	ID          string `db:"id" json:"id"`
	Emoji       string `db:"emoji" json:"emoji"`
	Description string `db:"description" json:"description"`
}
