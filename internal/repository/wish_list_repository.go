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
	newWishList := new(model.WishList)
	query := `
		INSERT INTO wish_lists (user_id, name, visibility)
		VALUES ($1, $2, $3)
		RETURNING *;
	`

	err := r.db.Get(
		newWishList, query, wishList.UserId, wishList.Name, wishList.Visibility,
	)
	if err != nil {
		// TODO: log error
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to create the wish list",
		)
	}

	return newWishList, nil
}

// Gets a wish list by its id.
func (r *WishListRepository) GetWishListById(
	userId int, wishListId int,
) (*model.WishList, *exception.ApiError) {
	wishList := new(model.WishList)
	query := `
		SELECT *
		FROM wish_lists
		WHERE user_id = $1 AND id = $2
	`

	err := r.db.Get(wishList, query, userId, wishListId)
	if err != nil {
		// TODO: log error
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to fetch the wish list",
		)
	}

	return wishList, nil
}

// Gets all the wish lists.
func (r *WishListRepository) ListWishLists(
	userId int,
) ([]*model.WishList, *exception.ApiError) {
	wishLists := []*model.WishList{}
	query := `
		SELECT *
		FROM wish_lists
		WHERE user_id = $1
	`

	err := r.db.Select(&wishLists, query, userId)
	if err != nil {
		// TODO: log error
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to fetch the wish lists",
		)
	}

	return wishLists, nil
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
