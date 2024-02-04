package config

import (
	"os"

	"github.com/brunomc/api-go/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	//create db and connection
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("database is not found, creating...")
		//create db file
		err = os.Mkdir("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Create file db failed %v", err)
			return nil, err
		}
		file.Close()
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("SQL error %v", err)
		return nil, err
	}
	//migrete the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("SQLITE automigrate error %v", err)
		return nil, err
	}
	return db, nil
}
