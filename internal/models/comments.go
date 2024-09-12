package models

import "time"

type Comment struct {
	ID              string    `db:"id" json:"id"`
	UserID          string    `db:"user_id" json:"user_id"`
	PostID          string    `db:"post_id" json:"post_id"`
	ParentID        string    `db:"parent_id" json:"parent_id"`
	Content         string    `db:"content" json:"content"`
	Timestamp       time.Time `db:"datetime" json:"datetime"`
	SubcommentCount int       `db:"subcomment_count" json:"subcomment_count"`
	Nickname        string    `json:"nickname"`
}
