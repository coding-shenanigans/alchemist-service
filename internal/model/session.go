package model

import "time"

type Session struct {
	Id           int
	UserId       int       `db:"user_id"`
	RefreshToken string    `db:"refresh_token"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
