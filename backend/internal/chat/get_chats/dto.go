package getchats

import "time"

type Input struct {
	UserID        int
	LimitMessages int
}

type Output struct {
	Chats []OutputChat `json:"chats"`
}

type OutputChat struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Users    []OutputUser     `json:"users"`
	Messages []OutputMessage  `json:"messages"`
}

type OutputUser struct {
	ID   int    `json:"id"`
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

type OutputMessage struct {
	UserAuthor OutputUser         `json:"user_author"`
	Message    OutputMessageContent `json:"message"`
}

type OutputMessageContent struct {
	ID   int       `json:"id"`
	Text string    `json:"text"`
	Time time.Time `json:"time"`
}
