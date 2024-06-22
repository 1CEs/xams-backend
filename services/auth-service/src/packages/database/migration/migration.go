package migration

import "gorm.io/gorm"

func RunMigration(db *gorm.DB, models ...interface{}) error {
	for _, model := range models {
		if err := db.AutoMigrate(&model); err != nil {
			return err
		}
	}
	return nil
}