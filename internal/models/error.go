package models

import "errors"

var (
	// ErrNoRecordFound reps "not found" error case.
	ErrNoRecordFound = errors.New("record not found")
)
