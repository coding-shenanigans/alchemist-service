package model

import "time"

type Item struct {
	Id               int       `json:"id"`
	WishListId       int       `json:"wishListId" db:"wish_list_id"`
	Url              string    `json:"url"`
	Name             string    `json:"name"`
	Price            float64   `json:"price"`
	Status           string    `json:"status"`
	ReservedByUserId *int      `json:"reservedByUserId" db:"reserved_by_user_id"`
	CreatedAt        time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt        time.Time `json:"updatedAt" db:"updated_at"`
}
