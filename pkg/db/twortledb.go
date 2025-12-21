package db

import (
	"fmt"
	"path/filepath"
	"twortle/pkg/db/tables"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitSQLiteConnection initializes the SQLite database
func InitSQLiteConnection() (*gorm.DB, error) {
	databasePath := filepath.Join("./", "wordbase.db")

	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening SQLite database: %v", err)
	}
	return db, nil
}

// InitSQLiteTables initializes the SQLite database tables
func InitSQLiteTables(db *gorm.DB) error {
	// Migrate the schemas
	dbCreateError := db.AutoMigrate(&tables.Word{})
	if dbCreateError != nil {
		fmt.Printf("Error creating table: %v\n", dbCreateError)
	}

	return nil
}

// CloseConnection closes the database connection
func CloseConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("error getting database instance: %v", err)
	}
	return sqlDB.Close()
}
