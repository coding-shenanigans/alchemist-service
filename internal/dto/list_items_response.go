package dto

import "github.com/coding-shenanigans/alchemist-service/internal/model"

type ListItemsResponse struct {
	Items []*model.Item `json:"items"`
}
