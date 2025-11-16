package repository

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/coding-shenanigans/alchemist-service/internal/exception"
	"github.com/coding-shenanigans/alchemist-service/internal/model"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Returns an error if the email already exists.
func (r *UserRepository) EmailExists(email string) *exception.ApiError {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE LOWER(email) = LOWER($1)
		);
	`

	err := r.db.Get(&exists, query, email)
	if err != nil {
		// TODO: log error
		return exception.NewApiError(
			http.StatusInternalServerError, "failed to check if email exists",
		)
	}

	if exists {
		return exception.NewApiError(
			http.StatusConflict, fmt.Sprintf("the email %q already exists", email),
		)
	}

	return nil
}

// Returns an error if the username already exists.
func (r *UserRepository) UsernameExists(username string) *exception.ApiError {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE LOWER(username) = LOWER($1)
		);
	`

	err := r.db.Get(&exists, query, username)
	if err != nil {
		// TODO: log error
		return exception.NewApiError(
			http.StatusInternalServerError, "failed to check if username exists",
		)
	}

	if exists {
		return exception.NewApiError(
			http.StatusConflict,
			fmt.Sprintf("the username %q already exists", username),
		)
	}

	return nil
}

// Creates a new user.
func (r *UserRepository) CreateUser(
	email string, username string, password string,
) (*model.User, *exception.ApiError) {
	var id int
	query := `
		INSERT INTO users (email, username, password)
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	err := r.db.Get(&id, query, email, username, password)
	if err != nil {
		// TODO: log error
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to create the user",
		)
	}

	user := new(model.User)
	query = `
		SELECT *
		FROM users
		WHERE id = $1
		LIMIT 1;
	`

	err = r.db.Get(user, query, id)
	if err != nil {
		// TODO: log error
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to fetch the user",
		)
	}

	return user, nil
}
