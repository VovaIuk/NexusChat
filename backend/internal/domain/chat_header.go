package domain

import "time"

type ChatHeader struct {
	ChatID      int
	ChatName    string
	LastMessage string
	LastTime    time.Time
}
