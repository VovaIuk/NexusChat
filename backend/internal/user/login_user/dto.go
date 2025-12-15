package login_user

type Input struct {
	Usertag  string `form:"tag"`
	Username string `form:"name"`
	Password string `form:"password"`
}

type Output struct {
	Id       int    `json:"id"`
	Usertag  string `json:"tag"`
	Username string `json:"name"`
	Password string `json:"password"`
}
