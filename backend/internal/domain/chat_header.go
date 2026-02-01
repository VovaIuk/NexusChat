package domain

import "time"

type ChatHeader struct {
	ChatID      int
	ChatName    string
	LastMessage string
	LastTime    time.Time
}

type ChatMember struct {
	ChatID   int
	UserID   int
	Usertag  string
	Username string
}

type ChatMessage struct {
	ChatID    int
	UserID    int
	Usertag   string
	Username  string
	MessageID int
	Text      string
	Time      time.Time
}
