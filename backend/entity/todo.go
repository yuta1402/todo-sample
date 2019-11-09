package entity

import (
	"time"
)

type Todo struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Completed bool       `json:"completed"`
	Created   *time.Time `json:"created"`
	Updated   *time.Time `json:"updated"`
}
