package register_user

type Input struct {
	Usertag  string `json:"tag"`
	Username string `json:"name"`
	Password string `json:"password"`
}

type Output struct {
	Id       int    `json:"id"`
	Usertag  string `json:"tag"`
	Username string `json:"name"`
}
