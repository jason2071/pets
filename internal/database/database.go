package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect opens a GORM connection to PostgreSQL using the given DSN.
func Connect(dsn string) (*gorm.DB, error) {
	// TranslateError maps driver errors to GORM sentinels (e.g. ErrDuplicatedKey),
	// so callers can detect unique violations without parsing SQLSTATE strings.
	return gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
}
