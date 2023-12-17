package model

import "time"

type Post struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Time    time.Time `json:"time"`
	Content string    `json:"content"`
}
