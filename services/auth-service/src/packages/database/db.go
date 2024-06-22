package database

import (
	"errors"

	"github.com/xams-backend/services/auth-service/config"
	"github.com/xams-backend/services/auth-service/src/internal/models"
	"github.com/xams-backend/services/auth-service/src/packages/database/migration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	Database struct {
		DB *gorm.DB
	}

	IDatabase interface {
		Connect() (*gorm.DB, error)
		AutoMigration() error
	}
)

func NewDatabase() IDatabase {
	return &Database{
		DB: nil,
	}
}

func (db *Database) Connect() (*gorm.DB, error) {
	config := config.NewConfig()
	dsn := config.DatabaseConfig()
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.DB = connection
	return db.DB, nil
}

func (db *Database) AutoMigration() error {
	if db.DB == nil {
		return errors.New("database connection is not initialized")
	}
	return migration.RunMigration(db.DB, &models.Faculty{}, &models.Branch{}, &models.User{})
}
