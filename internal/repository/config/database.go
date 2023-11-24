package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"profile-sandbox/internal/model/sandbox"
)

var DB = NewConnection()

func NewConnection() *gorm.DB {
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("Unabled to connect to database")
	}

	err = db.AutoMigrate(&sandbox.Scope{})
	if err != nil {
		panic("Failed to migrate Sandbox: " + err.Error())
	}
	return db
}
