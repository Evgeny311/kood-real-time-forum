package models

import "time"

type Chat struct {
	ID    string `db:"id" json:"id"`
	Users []ChatUser
	Lines []Line
}

type ChatUser struct {
	ID       string `db:"id" json:"id"`
	UserID   string `db:"user_id" json:"user_id"`
	ChatID   string `db:"chat_id" json:"chat_id"`
	Username string `db:"username" json:"username"`
}

type Line struct {
	ID        string    `db:"id" json:"id"`
	ChatID    string    `db:"chat_id" json:"chat_id"`
	UserID    string    `db:"user_id" json:"user_id"`
	Content   string    `db:"content" json:"content"`
	TimeStamp time.Time `db:"timestamp" json:"timestamp"`
}
