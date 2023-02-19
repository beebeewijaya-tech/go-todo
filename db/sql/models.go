package sql

import "time"

type PriorityType int32

const (
	LOW    PriorityType = 1
	MEDIUM PriorityType = 2
	HIGH   PriorityType = 3
)

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Todo struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Priority    PriorityType `json:"priority"`
	Author      int64        `json:"author"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
