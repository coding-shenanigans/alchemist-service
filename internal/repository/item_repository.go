package repository

import (
	"database/sql"
	"errors"
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

// Gets an item by its id.
func (r *ItemRepository) GetItemById(
	wishListId int, itemId int,
) (*model.Item, *exception.ApiError) {
	item := new(model.Item)
	query := `
		SELECT *
		FROM items
		WHERE wish_list_id = $1 AND id = $2
	`

	err := r.db.Get(item, query, wishListId, itemId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exception.NewApiError(
				http.StatusNotFound, "the item was not found",
			)
		} else {
			// TODO: log error
			return nil, exception.NewApiError(
				http.StatusInternalServerError, "failed to fetch the item",
			)
		}
	}

	return item, nil
}

// Gets all the items.
func (r *ItemRepository) ListItems(
	wishListId int,
) ([]*model.Item, *exception.ApiError) {
	items := []*model.Item{}
	query := `
		SELECT *
		FROM items
		WHERE wish_list_id = $1
		ORDER BY updated_at DESC
	`

	err := r.db.Select(&items, query, wishListId)
	if err != nil {
		// TODO: log error
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to fetch the items",
		)
	}

	return items, nil
}

// Updates an item.
func (r *ItemRepository) UpdateItem(
	item *model.Item,
) (*model.Item, *exception.ApiError) {
	updatedItem := new(model.Item)
	query := `
		UPDATE items
		SET 
			url = $1,
			name = $2,
			price = $3
		WHERE wish_list_id = $4 AND id = $5
		RETURNING *;
	`

	err := r.db.Get(
		updatedItem,
		query,
		item.Url,
		item.Name,
		item.Price,
		item.WishListId,
		item.Id,
	)
	if err != nil {
		// TODO: log error
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to update the item",
		)
	}

	return updatedItem, nil
}

// Deletes an item by its id.
func (r *ItemRepository) DeleteItemById(
	wishListId int, itemId int,
) *exception.ApiError {
	query := `
		DELETE FROM items
		WHERE wish_list_id = $1 AND id = $2;
	`

	_, err := r.db.Exec(query, wishListId, itemId)
	if err != nil {
		// TODO: log error
		return exception.NewApiError(
			http.StatusInternalServerError, "failed to delete the item",
		)
	}

	return nil
}
