package validator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name        string
		email       string
		expectError bool
	}{
		{
			name:        "blank",
			email:       "",
			expectError: true},
		{
			name:        "missing_local_part",
			email:       "@konoha.gov",
			expectError: true},
		{
			name:        "missing_at_sign",
			email:       "rock.lee.konoha.gov",
			expectError: true},
		{
			name:        "missing_domain",
			email:       "rock.lee@",
			expectError: true},
		{
			name:        "exceed_max_langth",
			email:       strings.Repeat("a", 255) + "@konoha.gov",
			expectError: true},
		{
			name:        "valid",
			email:       "rock.lee@konoha.gov",
			expectError: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.email)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		name        string
		username    string
		expectError bool
	}{
		{
			name:        "blank",
			username:    "",
			expectError: true,
		},
		{
			name:        "below_min_length",
			username:    "rl",
			expectError: true,
		},
		{
			name:        "exceed_max_length",
			username:    strings.Repeat("a", 37),
			expectError: true,
		},
		{
			name:        "invalid_characters",
			username:    "rock.lee",
			expectError: true,
		},
		{
			name:        "valid",
			username:    "rock_lee",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUsername(tt.username)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name        string
		password    string
		expectError bool
	}{
		{
			name:        "blank",
			password:    "",
			expectError: true,
		},
		{
			name:        "below_min_length",
			password:    "youth",
			expectError: true,
		},
		{
			name:        "exceed_max_length",
			password:    strings.Repeat("aA1.", 73),
			expectError: true,
		},
		{
			name:        "missing_lowercase_character",
			password:    "POW3R.OF.YOUTH",
			expectError: true,
		},
		{
			name:        "missing_uppercase_character",
			password:    "pow3r.of.youth",
			expectError: true,
		},
		{
			name:        "missing_number",
			password:    "power.of.YOUTH",
			expectError: true,
		},
		{
			name:        "missing_special_character",
			password:    "pow3rofYOUTH",
			expectError: true,
		},
		{
			name:        "valid",
			password:    "pow3r.of.YOUTH",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePassword(tt.password)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateWishListName(t *testing.T) {
	tests := []struct {
		name         string
		wishListName string
		expectError  bool
	}{
		{
			name:         "below_min_length",
			wishListName: "",
			expectError:  true,
		},
		{
			name:         "exceed_max_length",
			wishListName: strings.Repeat("a", 101),
			expectError:  true,
		},
		{
			name:         "valid",
			wishListName: "Splendid Ninja Training",
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateWishListName(tt.wishListName)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateWishListVisibility(t *testing.T) {
	tests := []struct {
		name        string
		visibility  string
		expectError bool
	}{
		{
			name:        "blank",
			visibility:  "",
			expectError: true,
		},
		{
			name:        "invalid_visibility",
			visibility:  "unknown",
			expectError: true,
		},
		{
			name:        "valid",
			visibility:  "friends_only",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateWishListVisibility(tt.visibility)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateUrl(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		expectError bool
	}{
		{
			name:        "blank",
			url:         "",
			expectError: true,
		},
		{
			name:        "exceed_max_length",
			url:         strings.Repeat("a", 2049),
			expectError: true,
		},
		{
			name:        "valid",
			url:         "example.com",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUrl(tt.url)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateItemName(t *testing.T) {
	tests := []struct {
		name        string
		itemName    string
		expectError bool
	}{
		{
			name:        "blank",
			itemName:    "",
			expectError: true,
		},
		{
			name:        "exceed_max_length",
			itemName:    strings.Repeat("a", 101),
			expectError: true,
		},
		{
			name:        "valid",
			itemName:    "my item name",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateItemName(tt.itemName)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
