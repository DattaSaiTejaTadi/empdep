package store

import (
	"context"

	"github.com/LetsFocus/account-service/empdep/models"
)

func (s *store) CreateEmployee(ctx context.Context, employee models.Employee) (models.Employee, error) {
	query := "INSERT INTO employee (id, name, dob, major, department) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, dob, major, department"
	row := s.db.QueryRowContext(ctx, query, employee.ID, employee.Name, employee.DOB, employee.Major, employee.Department)
	var createdEmployee models.Employee
	err := row.Scan(&createdEmployee.ID, &createdEmployee.Name, &createdEmployee.DOB, &createdEmployee.Major, &createdEmployee.Department)
	if err != nil {
		return models.Employee{}, err
	}
	return createdEmployee, nil
}
func (s *store) UpdateEmployee(ctx context.Context, employee models.Employee) (models.Employee, error) {
	query := "UPDATE employee SET name = $1, dob = $2, major = $3, department = $4 WHERE id = $5 RETURNING id, name, dob, major, department"
	row := s.db.QueryRowContext(ctx, query, employee.Name, employee.DOB, employee.Major, employee.Department, employee.ID)
	var updatedEmployee models.Employee
	err := row.Scan(&updatedEmployee.ID, &updatedEmployee.Name, &updatedEmployee.DOB, &updatedEmployee.Major, &updatedEmployee.Department)
	if err != nil {
		return models.Employee{}, err
	}
	return updatedEmployee, nil
}
func (s *store) DeleteEmployee(ctx context.Context, employee models.Employee) error {
	query := "DELETE FROM employee WHERE id = $1"
	_, err := s.db.ExecContext(ctx, query, employee.ID)
	return err
}
func (s *store) GetEmployee(ctx context.Context) ([]models.Employee, error) {
	query := "SELECT id, name, dob, major, department FROM employee"
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var employee models.Employee
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.DOB, &employee.Major, &employee.Department); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}
