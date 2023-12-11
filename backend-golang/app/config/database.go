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
	db.Set("gorm:auto_preload", true)

	log.Printf("Connection to the database was successful")
	return db
}
