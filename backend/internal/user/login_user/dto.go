package login_user

type Input struct {
	Usertag  string `json:"tag"`
	Password string `json:"password"`
}

type Output struct {
	User  User  `json:"user"`
	Token Token `json:"token"`
}

type User struct {
	ID   int    `json:"id"`
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

type Token struct {
	Refresh   string `json:"refresh"`
	ExpiresIn int    `json:"expiresIn"`
	Type      string `json:"type"`
}
