package search_user

//TODO: установить ограничение длина тега больше 3
type Input struct {
	Tag   string
	Limit int
}

type Output struct {
	Users []OutputUser `json:"users"`
}

type OutputUser struct {
	ID   int    `json:"id"`
	Tag  string `json:"tag"`
	Name string `json:"name"`
}
