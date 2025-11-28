package repository

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/coding-shenanigans/alchemist-service/internal/exception"
	"github.com/coding-shenanigans/alchemist-service/internal/model"
)

type SessionRepository struct {
	db *sqlx.DB
}

func NewSessionRepository(db *sqlx.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

// Creates a new session.
func (r *SessionRepository) CreateSession(
	userId int, refreshToken string,
) (*model.Session, *exception.ApiError) {
	var id int
	query := `
		INSERT INTO sessions (user_id, refresh_token)
		VALUES ($1, $2)
		RETURNING id;
	`

	err := r.db.Get(&id, query, userId, refreshToken)
	if err != nil {
		// TODO: log error
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to create the session",
		)
	}

	session := new(model.Session)
	query = `
		SELECT *
		FROM sessions
		WHERE id = $1
		LIMIT 1;
	`

	err = r.db.Get(session, query, id)
	if err != nil {
		// TODO: log error
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to fetch the session",
		)
	}

	return session, nil
}

// Gets a session by its refresh token.
func (r *SessionRepository) GetSessionByRefreshToken(
	refreshToken string,
) (*model.Session, *exception.ApiError) {
	session := new(model.Session)
	query := `
		SELECT *
		FROM sessions
		WHERE refresh_token = $1
		LIMIT 1;
	`

	err := r.db.Get(session, query, refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exception.NewApiError(
				http.StatusNotFound, "the session was not found",
			)
		} else {
			// TODO: log error
			return nil, exception.NewApiError(
				http.StatusInternalServerError, "failed to fetch the session",
			)
		}
	}

	return session, nil
}

// Refreshes a session by updating its refresh token.
func (r *SessionRepository) RefreshSession(
	id int, refreshToken string,
) *exception.ApiError {
	query := `
		UPDATE sessions
		SET refresh_token = $1
		WHERE id = $2;
	`

	_, err := r.db.Exec(query, refreshToken, id)
	if err != nil {
		// TODO: log error
		return exception.NewApiError(
			http.StatusInternalServerError, "failed to refresh the session",
		)
	}

	return nil
}

// Deletes a session by its refresh token.
func (r *SessionRepository) DeleteSessionByRefreshToken(
	refreshToken string,
) *exception.ApiError {
	query := `
		DELETE FROM sessions
		WHERE refresh_token = $1;
	`

	_, err := r.db.Exec(query, refreshToken)
	if err != nil {
		// TODO: log error
		return exception.NewApiError(
			http.StatusInternalServerError, "failed to delete the session",
		)
	}

	return nil
}
