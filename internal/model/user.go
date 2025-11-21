package model

import "time"

type User struct {
	Id        int
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
