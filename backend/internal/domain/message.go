package domain

import "time"

type Message struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	ChatID int       `json:"chat_id"`
	Text   string    `json:"text"`
	Time   time.Time `json:"time"`
}
