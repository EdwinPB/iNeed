package models

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

// Client model struct.
type Client struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Clients array model struct of Client.
type Clients []Client

// String converts the struct into a string value.
func (c Client) String() string {
	return fmt.Sprintf("%+v\n", c)
}
