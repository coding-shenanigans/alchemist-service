package validator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	testCases := []struct {
		name        string
		email       string
		expectError bool
	}{
		{"Blank", "", true},
		{"MissingLocalPart", "@konoha.gov", true},
		{"MissingAtSign", "rock.lee.konoha.gov", true},
		{"MissingDomain", "rock.lee@", true},
		{"ExceedMaxLength", strings.Repeat("a", 255) + "@konoha.gov", true},
		{"Valid", "rock.lee@konoha.gov", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateEmail(tc.email)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateUsername(t *testing.T) {
	testCases := []struct {
		name        string
		username    string
		expectError bool
	}{
		{"Blank", "", true},
		{"BelowMinLength", "rl", true},
		{"ExceedMaxLength", strings.Repeat("a", 37), true},
		{"InvalidCharacters", "rock.lee", true},
		{"Valid", "rock_lee", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateUsername(tc.username)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	testCases := []struct {
		name        string
		password    string
		expectError bool
	}{
		{"Blank", "", true},
		{"BelowMinLength", "youth", true},
		{"ExceedMaxLength", strings.Repeat("aA1.", 73), true},
		{"MissingLowercaseCharacter", "POW3R.OF.YOUTH", true},
		{"MissingUppercaseCharacter", "pow3r.of.youth", true},
		{"MissingNumber", "power.of.YOUTH", true},
		{"MissingSpecialCharacter", "pow3rofYOUTH", true},
		{"Valid", "pow3r.of.YOUTH", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidatePassword(tc.password)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
