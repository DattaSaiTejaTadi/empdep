package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/LetsFocus/account-service/empdep/models"
	"github.com/google/uuid"
)

func (a *app) CreateEmployee(ctx context.Context, employee models.Employee) (models.Employee, error) {
	employee.ID = uuid.New()
	var err error
	age := calculateAge(employee.DOB)
	if age < 22 {
		return models.Employee{}, errors.New("Age error, age must be >=22")
	}
	employee.Department, err = a.store.GetDepartmentID(strings.ToUpper(employee.Major))
	if err != nil {
		return models.Employee{}, err
	}
	return a.store.CreateEmployee(ctx, employee)
}
func (a *app) UpdateEmployee(ctx context.Context, employee models.Employee) (models.Employee, error) {
	return a.store.UpdateEmployee(ctx, employee)
}
func (a *app) DeleteEmployee(ctx context.Context, employee models.Employee) error {
	return a.store.DeleteEmployee(ctx, employee)
}
func (a *app) GetEmployee(ctx context.Context) ([]models.Employee, error) {
	return a.store.GetEmployee(ctx)
}

func calculateAge(dob time.Time) int {
	today := time.Now()
	age := today.Year() - dob.Year()

	if today.YearDay() < dob.YearDay() {
		age--
	}

	return age
}
