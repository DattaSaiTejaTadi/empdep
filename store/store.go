package store

import (
	"database/sql"
	"log"
	"time"

	"github.com/LetsFocus/account-service/empdep/metrics"
	_ "github.com/lib/pq"
)

type store struct {
	db *sql.DB
}

func New() *store {
	db, err := sql.Open("postgres", "postgresql://postgres:example@localhost/empdep?sslmode=disable")
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		metrics.FailedConnections.Inc() // Increment the failed connections counter
		panic(err)
	}

	// Set database connection parameters
	db.SetMaxOpenConns(10) // Set the maximum number of open connections
	db.SetMaxIdleConns(5)  // Set the maximum number of idle connections

	// Start a goroutine to collect connection metrics
	go func() {
		for {
			stats := db.Stats()
			metrics.ActiveConnections.Set(float64(stats.InUse))
			metrics.IdleConnections.Set(float64(stats.Idle))
			metrics.TotalConnections.Set(float64(stats.OpenConnections))

			// Collect metrics every 10 seconds
			time.Sleep(10 * time.Second)
		}
	}()

	return &store{
		db: db,
	}
}
