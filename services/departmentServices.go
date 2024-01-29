package services

import (
	"context"

	"github.com/LetsFocus/account-service/empdep/models"
)

func (a *app) GetDepatments(ctx context.Context) ([]models.Department, error) {
	return a.store.GetDepatments(ctx)
}
func (a *app) CreateDepartment(ctx context.Context, department models.Department) (models.Department, error) {
	return a.store.CreateDepartment(ctx, department)
}
func (a *app) UpdateDepartment(ctx context.Context, department models.Department) (models.Department, error) {
	return a.store.UpdateDepartment(ctx, department)
}
func (a *app) DeleteDepartment(ctx context.Context, department models.Department) error {
	return a.store.DeleteDepartment(ctx, department)
}
