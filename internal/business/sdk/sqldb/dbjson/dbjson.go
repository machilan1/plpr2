// Package dbjson provides a type that can be used to convert JSON data to and from a SQL database.
package dbjson

import (
	"database/sql/driver"
	"encoding/json"
)

// JSONColumn is a type that can be used to store JSON data in a SQL database.
// Note that the struct tag for the field must be set to `json:"column_name"`.
type JSONColumn[T any] struct {
	v *T
}

// Scan implements the sql.Scanner interface.
func (j *JSONColumn[T]) Scan(src any) error {
	if src == nil {
		j.v = nil
		return nil
	}
	j.v = new(T)
	return json.Unmarshal(src.([]byte), j.v)
}

// Value implements the driver.Valuer interface.
// Must NOT be a pointer receiver, otherwise it won't be called.
func (j JSONColumn[T]) Value() (driver.Value, error) {
	raw, err := json.Marshal(j.v)
	return raw, err
}

// Get returns the value of the JSONColumn.
func (j *JSONColumn[T]) Get() T {
	if j.v == nil {
		v := new(T)
		return *v
	}
	return *j.v
}

func (j *JSONColumn[T]) Set(v T) {
	j.v = &v
}
