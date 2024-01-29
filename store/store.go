package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type store struct {
	db *sql.DB
}

func New() *store {
	db, err := sql.Open("postgres", "postgresql://postgres:password@localhost/empDep?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return &store{
		db: db,
	}
}
