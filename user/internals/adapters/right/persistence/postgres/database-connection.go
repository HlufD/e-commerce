package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseConnectionAdapter struct {
	dataSource string
	driverName string
}

func NewDatabaseConnectionAdapter(dataSource string, driverName string) *DatabaseConnectionAdapter {
	return &DatabaseConnectionAdapter{dataSource, driverName}
}

func (d *DatabaseConnectionAdapter) Connect() (*sql.DB, error) {
	db, err := sql.Open(d.driverName, d.dataSource)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	return db, nil
}

func (d *DatabaseConnectionAdapter) Close(db *sql.DB) error {
	if db != nil {
		return db.Close()
	}
	return nil
}
