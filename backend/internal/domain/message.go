package domain

import "time"

type Message struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	ChatID int       `json:"chat_id"`
	Text   string    `json:"text"`
	Time   time.Time `json:"time"`
}

type MessageCreate struct {
	UserID int
	ChatID int
	Text   string
	Time   time.Time
}

func NewMessageCreate(userID, chatID int, text string, time time.Time) MessageCreate {
	return MessageCreate{
		UserID: userID,
		ChatID: chatID,
		Text:   text,
		Time:   time,
	}
}
