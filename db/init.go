package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	// synchronous updates
	syncResult := db.Exec("PRAGMA synchronous = NORMAL")
	if syncResult.Error != nil {
		return nil, syncResult.Error
	}

	err = db.AutoMigrate()
	if err != nil {
		return nil, err
	}
	return db, nil
}
