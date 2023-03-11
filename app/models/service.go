package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

// Service model struct.
type Service struct {
	ID         uuid.UUID `json:"id" db:"id"`
	BusinessID uuid.UUID `json:"business_id" db:"business_id"`
	ClientID   uuid.UUID `json:"client_id" db:"client_id"`
	Status     string    `json:"status" db:"status"`

	Business  Business  `belongs_to:"businesses" fk_id:"BusinessID"`
	Client    Client    `belongs_to:"clients" fk_id:"ClientID"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Services array model struct of Service.
type Services []Service

// String converts the struct into a string value.
func (s Service) String() string {
	jc, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("error marshalling json on string nethod: %v\n", err)
	}
	return string(jc)
}
