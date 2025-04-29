package services

import (
	"context"
	"strings"

	"github.com/LetsFocus/account-service/empdep/models"
	"github.com/google/uuid"
)

func (a *app) GetDepatments(ctx context.Context) ([]models.Department, error) {
	return a.store.GetDepatments(ctx)
}
func (a *app) CreateDepartment(ctx context.Context, department models.Department) (models.Department, error) {
	department.ID = uuid.New()
	department.Name = strings.ToUpper(department.Name)
	return a.store.CreateDepartment(ctx, department)
}
func (a *app) UpdateDepartment(ctx context.Context, department models.Department) (models.Department, error) {
	return a.store.UpdateDepartment(ctx, department)
}
func (a *app) DeleteDepartment(ctx context.Context, department models.Department) error {
	return a.store.DeleteDepartment(ctx, department)
}
