package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectToDb() *gorm.DB {

	LoadEnv()

	dsn := os.Getenv("DB_URI")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection failed")
	}
	log.Printf("Database connection successful")
	db.Set("gorm:auto_preload", true)

	return db
}
