package models

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	DOB        time.Time `json:"dob"`
	Major      string    `json:"major"`
	Department int       `json:"department"`
}
