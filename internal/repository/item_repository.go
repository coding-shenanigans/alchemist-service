package repository

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/coding-shenanigans/alchemist-service/internal/exception"
	"github.com/coding-shenanigans/alchemist-service/internal/model"
)

type ItemRepository struct {
	db *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

// Creates a new item.
func (r *ItemRepository) CreateItem(
	item *model.Item,
) (*model.Item, *exception.ApiError) {
	newItem := new(model.Item)
	query := `
		INSERT INTO items (wish_list_id, url, name, price)
		VALUES ($1, $2, $3, $4)
		RETURNING *;
	`

	err := r.db.Get(
		newItem, query, item.WishListId, item.Url, item.Name, item.Price,
	)
	if err != nil {
		// TODO: log error
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to add the item",
		)
	}

	return newItem, nil
}
