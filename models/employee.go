package models

import (
	"github.com/google/uuid"
)

type Employee struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	DOB        string    `json:"dob"`
	Major      string    `json:"major"`
	Department uuid.UUID `json:"department"`
	DepName    string    `json:"departmentName"`
}

type Employees struct {
	EmployeesCount int        `json:"EmployeesCount"`
	EmployeesData  []Employee `json:"EmployeesData"`
}
