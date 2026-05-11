package model

import "time"

type Session struct {
	Id             int
	UserId         int       `db:"user_id"`
	RefreshTokenId string    `db:"refresh_token_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
