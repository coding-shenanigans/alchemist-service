package model

import "time"

type WishList struct {
	Id         int       `json:"id"`
	UserId     int       `json:"userId" db:"user_id"`
	Name       string    `json:"name"`
	Visibility string    `json:"visibility"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
}
