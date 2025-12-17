package get_chatheaders

type Input struct {
	JWTToken string `json:"-"`
}

type OutputChat struct {
	ID      int    `json:"chat_id"`
	Name    string `json:"chat_name"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

type Output struct {
	Chats []OutputChat `json:"chats"`
}
