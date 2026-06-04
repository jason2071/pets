package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect opens a GORM connection to PostgreSQL using the given DSN.
func Connect(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
