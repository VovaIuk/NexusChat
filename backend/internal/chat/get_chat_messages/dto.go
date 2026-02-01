package get_chat_messages

import "time"

type Input struct {
	UserID int // jwt token для провекри принадлежности к чату
	ChatID int // url path param
	Limit  int
	Offset int
}

type Output struct {
	Messages []OutputMessage `json:"messages"`
}

type OutputUser struct {
	ID   int    `json:"id"`
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

type OutputMessage struct {
	UserAuthor OutputUser           `json:"user_author"`
	Message    OutputMessageContent `json:"message"`
}

type OutputMessageContent struct {
	ID   int       `json:"id"`
	Text string    `json:"text"`
	Time time.Time `json:"time"`
}
