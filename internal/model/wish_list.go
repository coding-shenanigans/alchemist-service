package model

import "time"

type WishList struct {
	Id         int
	UserId     int `db:"user_id"`
	Name       string
	Visibility string
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
