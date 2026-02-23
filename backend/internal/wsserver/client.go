package wsserver

import "github.com/gorilla/websocket"

type Client struct {
	UserID   int
	Usertag  string
	Username string
	Rooms    []int

	Conn *websocket.Conn
}

func NewDefaultClient(conn *websocket.Conn) *Client {
	return &Client{
		UserID:   0,
		Usertag:  "",
		Username: "",
		Rooms:    make([]int, 0),
		Conn:     conn,
	}
}
