package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDB(dbUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{}) //Config should ensure that settings like re-creation of tables etc.
	if err != nil {
		return nil, err
	}

	return db, err

}
