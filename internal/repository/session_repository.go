package repository

import (
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
		WHERE id = $1;
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
