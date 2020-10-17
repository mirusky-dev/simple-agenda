package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New ...
func New(connStr string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	return db, err
}

// Clone ...
func Clone(gormDB *gorm.DB) (*gorm.DB, error) {
	sqlConn, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	cloneDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlConn,
	}), &gorm.Config{})

	return cloneDB, err
}
