package models

import (
	"time"
)

type Post struct {
	ID           string            `db:"id" json:"id"`
	Title        string            `db:"title" json:"title"`
	Content      string            `db:"content" json:"content"`
	UserID       string            `db:"user_id" json:"user_id"`
	UserNickname string            `db:"nickname" json:"nickname"`
	ImageURL     string            `db:"image_url" json:"image_url"`
	Timestamp    time.Time         `db:"timestamp" json:"timestamp"`
	CommentCount int               `db:"comment_count" json:"comment_count"`
}
