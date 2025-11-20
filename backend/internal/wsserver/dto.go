package wsserver

type wsMessage struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Time     string `json:"time"`
}
