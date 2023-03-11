package models

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

// Service model struct.
type Service struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Services array model struct of Service.
type Services []Service

// String converts the struct into a string value.
func (s Service) String() string {
	return fmt.Sprintf("%+v\n", s)
}
