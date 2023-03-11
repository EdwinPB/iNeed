package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

// Business model struct.
type Business struct {
	ID uuid.UUID `json:"id" db:"id"`

	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Category    string   `json:"category" db:"category"`
	Phone       string   `json:"phone" db:"phone"`
	ServiceTime string   `json:"service_time" db:"service_time"`
	Address     string   `json:"address" db:"address"`
	Services    Services `json:"services" has_many:"services"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Businesses array model struct of Business.
type Businesses []Business

// String converts the struct into a string value.
func (b Business) String() string {
	jc, err := json.Marshal(b)
	if err != nil {
		fmt.Printf("error marshalling json on string nethod: %v\n", err)
	}
	return string(jc)
}
