package validator

import (
	"fmt"
	"net/mail"
	"regexp"
)

const (
	emailMaxLength             = 254
	usernameMinLength          = 3
	usernameMaxLength          = 36
	usernameRegexString        = "^\\w+$" // alphanumeric characters
	passwordMinLength          = 8
	passwordMaxLength          = 72
	passwordLowerRegexString   = "[a-z]{1}"        // 1+ lowercase characters
	passwordUpperRegexString   = "[A-Z]{1}"        // 1+ uppercase characters
	passwordNumberRegexString  = "[0-9]{1}"        // 1+ numbers
	passwordSpecialRegexString = "[^a-zA-Z0-9]{1}" // 1+ special characters
)

var (
	usernameRegex        = regexp.MustCompile(usernameRegexString)
	passwordLowerRegex   = regexp.MustCompile(passwordLowerRegexString)
	passwordUpperRegex   = regexp.MustCompile(passwordUpperRegexString)
	passwordNumberRegex  = regexp.MustCompile(passwordNumberRegexString)
	passwordSpecialRegex = regexp.MustCompile(passwordSpecialRegexString)
)

// Checks if the email is valid.
func ValidateEmail(email string) error {
	if len(email) > emailMaxLength {
		return fmt.Errorf(
			"the email address should not exceed %d characters", emailMaxLength,
		)
	}

	if _, err := mail.ParseAddress(email); err != nil {
		// TODO: log error
		return fmt.Errorf("the email address is not valid")
	}

	return nil
}

// Checks if the username is valid.
func ValidateUsername(username string) error {
	if len(username) < usernameMinLength {
		return fmt.Errorf(
			"the username should have at least %d characters", usernameMinLength,
		)
	}

	if len(username) > usernameMaxLength {
		return fmt.Errorf(
			"the username should not exceed %d characters", usernameMaxLength,
		)
	}

	if !usernameRegex.MatchString(username) {
		return fmt.Errorf(
			"the username should only contain alphanumeric characters",
		)
	}

	return nil
}

// Checks if the password is valid.
func ValidatePassword(password string) error {
	if len(password) < passwordMinLength {
		return fmt.Errorf(
			"the password should have at least %d characters", passwordMinLength,
		)
	}

	if len(password) > passwordMaxLength {
		return fmt.Errorf(
			"the password should not exceed %d characters", passwordMaxLength,
		)
	}

	if !passwordLowerRegex.MatchString(password) {
		return fmt.Errorf(
			"the password should contain at least 1 lowercase character",
		)
	}

	if !passwordUpperRegex.MatchString(password) {
		return fmt.Errorf(
			"the password should contain at least 1 uppercase character",
		)
	}

	if !passwordNumberRegex.MatchString(password) {
		return fmt.Errorf("the password should contain at least 1 number")
	}

	if !passwordSpecialRegex.MatchString(password) {
		return fmt.Errorf(
			"the password should contain at least 1 special character",
		)
	}

	return nil
}
