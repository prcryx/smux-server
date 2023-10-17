package entities

import (
	"time"
)



type BlogEntity struct {
	BlogId    int
	Title     string
	Body      string
	CreatedAt time.Time
	AuthorId  int
	Author    Author
}
