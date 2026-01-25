package wsserver

type wsMessage struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Message  string `json:"message"`
	Time     string `json:"time"`
}
