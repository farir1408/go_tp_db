package models

//easyjson:json
type Status struct {
	ForumSum  int `json:"forum"`
	PostsSum  int `json:"post"`
	ThreadSum int `json:"thread"`
	UserSum   int `json:"user"`
}
