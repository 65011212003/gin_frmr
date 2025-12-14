package database

import (
	"log"

	"gin_frmr/internal/repository"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewSQLiteDB(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")

	// Auto migrate
	if err := db.AutoMigrate(&repository.UserModel{}); err != nil {
		return nil, err
	}

	log.Println("Database migrated successfully")

	return db, nil
}
