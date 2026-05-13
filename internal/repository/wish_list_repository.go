package repository

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/coding-shenanigans/alchemist-service/internal/exception"
	"github.com/coding-shenanigans/alchemist-service/internal/model"
)

type WishListRepository struct {
	db *sqlx.DB
}

func NewWishListRepository(db *sqlx.DB) *WishListRepository {
	return &WishListRepository{db: db}
}

// Creates a new wish list.
func (r *WishListRepository) CreateWishList(
	wishList *model.WishList,
) (*model.WishList, *exception.ApiError) {
	// TODO: Implement function.
	return nil, exception.NewApiError(
		http.StatusNotImplemented, "not implemented yet",
	)
}

// Gets a wish list by its id.
func (r *WishListRepository) GetWishListById(
	id int,
) (*model.WishList, *exception.ApiError) {
	// TODO: Implement function.
	return nil, exception.NewApiError(
		http.StatusNotImplemented, "not implemented yet",
	)
}

// Gets all the wish lists.
func (r *WishListRepository) ListWishLists(
	pageSize int, pageToken string, filter string,
) ([]*model.WishList, *exception.ApiError) {
	// TODO: Implement function.
	return nil, exception.NewApiError(
		http.StatusNotImplemented, "not implemented yet",
	)
}

// Updates a wish list.
func (r *WishListRepository) UpdateWishList(
	wishList *model.WishList, updateMask []string,
) (*model.WishList, *exception.ApiError) {
	// TODO: Implement function.
	return nil, exception.NewApiError(
		http.StatusNotImplemented, "not implemented yet",
	)
}

// Deletes a wish list by its id.
func (r *WishListRepository) DeleteWishListById(id int) *exception.ApiError {
	// TODO: Implement function.
	return exception.NewApiError(
		http.StatusNotImplemented, "not implemented yet",
	)
}
