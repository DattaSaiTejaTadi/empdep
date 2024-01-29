package store

import (
	"context"

	"github.com/LetsFocus/account-service/empdep/models"
)

type Store interface {
	GetDepatments(ctx context.Context) ([]models.Department, error)
	CreateDepartment(ctx context.Context, department models.Department) (models.Department, error)
	UpdateDepartment(ctx context.Context, department models.Department) (models.Department, error)
	DeleteDepartment(ctx context.Context, department models.Department) error
	CreateEmployee(ctx context.Context, employee models.Employee) (models.Employee, error)
	UpdateEmployee(ctx context.Context, employee models.Employee) (models.Employee, error)
	DeleteEmployee(ctx context.Context, employee models.Employee) error
	GetEmployee(ctx context.Context) ([]models.Employee, error)
	GetDepartmentID(major string) (int, error)
}
