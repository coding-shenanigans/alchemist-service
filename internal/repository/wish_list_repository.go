package repository

import (
	"database/sql"
	"errors"
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exception.NewApiError(
				http.StatusNotFound, "the wish list was not found",
			)
		} else {
			// TODO: log error
			return nil, exception.NewApiError(
				http.StatusInternalServerError, "failed to fetch the wish list",
			)
		}
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
		ORDER BY updated_at DESC
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
	userId int, wishList *model.WishList,
) (*model.WishList, *exception.ApiError) {
	updatedWishList := new(model.WishList)
	query := `
		UPDATE wish_lists
		SET 
			name = $1,
			visibility = $2
		WHERE user_id = $3 AND id = $4
		RETURNING *;
	`

	err := r.db.Get(
		updatedWishList,
		query,
		wishList.Name,
		wishList.Visibility,
		wishList.UserId,
		wishList.Id,
	)
	if err != nil {
		// TODO: log error
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to update the wish list",
		)
	}

	return updatedWishList, nil
}

// Deletes a wish list by its id.
func (r *WishListRepository) DeleteWishListById(
	userId int, wishListId int,
) *exception.ApiError {
	query := `
		DELETE FROM wish_lists
		WHERE user_id = $1 AND id = $2;
	`

	_, err := r.db.Exec(query, userId, wishListId)
	if err != nil {
		// TODO: log error
		return exception.NewApiError(
			http.StatusInternalServerError, "failed to delete the wish list",
		)
	}

	return nil
}
