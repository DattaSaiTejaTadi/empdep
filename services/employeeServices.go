package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/LetsFocus/account-service/empdep/models"
	"github.com/google/uuid"
)

func (a *app) CreateEmployee(ctx context.Context, employee models.Employee) (models.Employee, error) {
	employee.ID = uuid.New()
	var err error
	fmt.Println("about recieve department if")
	employee.Department, employee.DepName, err = a.store.GetDepartmentID(strings.ToUpper(employee.Major))
	if err != nil {
		return models.Employee{}, err
	}
	fmt.Println(employee)
	return a.store.CreateEmployee(ctx, employee)
}
func (a *app) UpdateEmployee(ctx context.Context, employee models.Employee) (models.Employee, error) {
	var err error
	employee.Department, employee.DepName, err = a.store.GetDepartmentID(strings.ToUpper(employee.Major))
	if err != nil {
		return models.Employee{}, err
	}
	return a.store.UpdateEmployee(ctx, employee)
}
func (a *app) DeleteEmployee(ctx context.Context, employee models.Employee) error {
	return a.store.DeleteEmployee(ctx, employee)
}
func (a *app) GetEmployee(ctx context.Context) (models.Employees, error) {
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
