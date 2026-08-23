package model

type ItemWithUser struct {
	Item
	ReservedByUsername *string `json:"reservedByUsername" db:"reserved_by_username"`
}
