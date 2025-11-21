package dto

import "github.com/coding-shenanigans/alchemist-service/internal/validator"

type SignupRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validates the request fields.
func (r *SignupRequest) Validate() error {
	if err := validator.ValidateEmail(r.Email); err != nil {
		return err
	}

	if err := validator.ValidateUsername(r.Username); err != nil {
		return err
	}

	if err := validator.ValidatePassword(r.Password); err != nil {
		return err
	}

	return nil
}
