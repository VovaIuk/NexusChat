package getme

type Input struct {
	Token string `header:"Authorization"`
}

type Output struct {
	Id       int    `json:"id"`
	Usertag  string `json:"tag"`
	Username string `json:"name"`
}
