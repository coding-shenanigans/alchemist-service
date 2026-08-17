package dto

import (
	"bytes"
	"encoding/json"
)

var nullBytes = []byte("null")

// NullableField represents a JSON field that distinguishes between three states:
//   - Omitted: Field was not present in the JSON payload (IsSet = false, Value = nil).
//   - Explicit null: Field was set to `null` in JSON (IsSet = true, Value = nil).
//   - Value provided: Field was set to a non-null value (IsSet = true, Value = &val).
//
// This struct is designed for request DTOs (such as PATCH updates) where omitting
// a field means "do not modify", while passing `null` means "clear the value".
type NullableField[T any] struct {
	Value *T
	IsSet bool
}

func (f *NullableField[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(bytes.TrimSpace(data), nullBytes) {
		f.Value = nil
		f.IsSet = true
		return nil
	}

	value := new(T)
	if err := json.Unmarshal(data, value); err != nil {
		return err
	}

	f.Value = value
	f.IsSet = true

	return nil
}
