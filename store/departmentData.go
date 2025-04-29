package store

import (
	"context"

	"github.com/LetsFocus/account-service/empdep/models"
	"github.com/google/uuid"
)

func (s *store) GetDepatments(ctx context.Context) ([]models.Department, error) {
	var departments []models.Department
	rows, err := s.db.Query("select id,name,floor from department")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var department models.Department
		if err := rows.Scan(&department.ID, &department.Name, &department.Floor); err != nil {
			return nil, err
		}
		departments = append(departments, department)
	}
	return departments, nil
}
func (s *store) CreateDepartment(ctx context.Context, department models.Department) (models.Department, error) {
	query := `
	INSERT INTO department (id,name, floor)
	VALUES ($1, $2, $3)
	RETURNING id, name, floor
`
	// Prepare the statement.
	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return models.Department{}, err
	}
	defer stmt.Close()
	// Execute the statement and scan the result.
	var createdDepartment models.Department
	err = stmt.QueryRowContext(ctx, department.ID, department.Name, department.Floor).Scan(
		&createdDepartment.ID, &createdDepartment.Name, &createdDepartment.Floor,
	)
	if err != nil {
		return models.Department{}, err
	}
	return createdDepartment, nil
}
func (s *store) UpdateDepartment(ctx context.Context, department models.Department) (models.Department, error) {
	// Define your SQL update query.
	query := `
	  UPDATE department
	  SET name = $2, floor = $3
	  WHERE id = $1
	  RETURNING id, name, floor
  `

	var updatedDepartment models.Department

	// Execute the SQL query.
	err := s.db.QueryRowContext(ctx, query, department.ID, department.Name, department.Floor).
		Scan(&updatedDepartment.ID, &updatedDepartment.Name, &updatedDepartment.Floor)

	if err != nil {
		return models.Department{}, err
	}

	return updatedDepartment, nil
}
func (s *store) DeleteDepartment(ctx context.Context, department models.Department) error {
	// Define your SQL delete query.
	query := "DELETE FROM department WHERE id = $1"

	// Execute the SQL query to delete the department.
	_, err := s.db.ExecContext(ctx, query, department.ID)

	if err != nil {
		return err
	}

	return nil
}
func (s *store) GetDepartmentID(major string) (uuid.UUID, string, error) {
	var Dept uuid.UUID
	var DepName string
	var dberr error
	statement := "select id, name from department where name=$1"
	stmt, err := s.db.Prepare(statement)
	if err != nil {
		return uuid.UUID{}, "", err
	}
	if major == "MBA" {
		dberr = stmt.QueryRow("HR").Scan(&Dept, &DepName)
	} else if major == "CSE" || major == "MCA" {
		dberr = stmt.QueryRow("TECH").Scan(&Dept, &DepName)
	} else if major == "B.COM" || major == "CA" {
		dberr = stmt.QueryRow("ACCOUNTS").Scan(&Dept, &DepName)
	}
	if dberr != nil {
		return uuid.UUID{}, "", dberr
	}
	return Dept, DepName, nil
}
