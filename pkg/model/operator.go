package model

import (
	"time"
)

// Operator represents an operator data.
type Operator struct {
	ID      string    `json:"id" db:"id"`
	Name    string    `json:"name" db:"name"`
	Valid   bool      `json:"valid" db:"valid"`
	Created time.Time `json:"created" db:"created"`
}
