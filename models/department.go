package models

import "github.com/google/uuid"

type Department struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Floor int       `json:"floor,omitempty"`
}
