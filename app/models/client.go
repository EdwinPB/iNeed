package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

// Client model struct.
type Client struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Phone     string    `json:"phone" has_many:"phone"`
	Address   string    `json:"address" has_many:"address"`
	Services  Services  `json:"services" has_many:"services"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Clients array model struct of Client.
type Clients []Client

// String converts the struct into a string value.
func (c Client) String() string {
	jc, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("error marshalling json on string nethod: %v\n", err)
	}
	return string(jc)
}
