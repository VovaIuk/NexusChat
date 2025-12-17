package get_chathistory

type Input struct {
	ChatID   int    `json:"-"`
	JWTToken string `json:"-"`
}

type OutputMessage struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Message  string `json:"message"`
	Time     string `json:"time"`
}

type Output struct {
	Messages []OutputMessage `json:"messages"`
}
