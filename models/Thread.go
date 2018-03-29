package models

import "time"

//easyjson:json
type Thread struct {
	ID      int        `json:"id"`
	Author  string     `json:"author"`
	Created *time.Time `json:"created"`
	ForumId string     `json:"forum"`
	Message string     `json:"message"`
	Slug    string     `json:"slug"`
	Title   string     `json:"title"`
	Votes   int        `json:"votes"`
}

//easyjson:json
type ThreadUpdate struct {
	Message string `json:"message"`
	Title   string `json:"title"`
}

//easyjson:json
type Threads []*Thread
