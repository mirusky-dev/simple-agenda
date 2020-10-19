package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres dialect
)

// New create a new connection based on connection string
func New(connStr string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", connStr)
	return db, err
}
