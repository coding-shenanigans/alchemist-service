package dto

import "github.com/coding-shenanigans/alchemist-service/internal/model"

type CancelItemReservationResponse struct {
	Item *model.ItemWithUser `json:"item"`
}
