package wsserver

import "time"

type Type string

const (
	TypeAuth        Type = "auth"
	TypeChatMessage Type = "chat_message"
)

type Message struct {
	Type Type   `json:"type"`
	Data string `json:"data"`
}

type AuthData struct {
	JWTToken string `json:"token"`
}

type ChatMessageData struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

type BroadcastMessage struct {
	MessageID int       `json:"message_id"`
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Usertag   string    `json:"usertag"`
	ChatID    int       `json:"chat_id"`
	Text      string    `json:"text"`
	Time      time.Time `json:"time"`
}
