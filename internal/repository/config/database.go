package config

import (
	"fmt"
	"os"
	"profile-sandbox/internal/model/sandbox"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB = NewConnection()

func NewConnection() *gorm.DB {
	host := byEnvOrDefault("DB_HOST", "localhost")
	password := byEnvOrDefault("DB_PASSWORD", "1234")
	port := byEnvOrDefault("DB_PORT", "5432")
	dbName := byEnvOrDefault("DB_NAME", "profile_sandbox")
	user := byEnvOrDefault("DB_USER", "postgres")

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

func byEnvOrDefault(env string, def string) string {
	value := os.Getenv(env)
	if value == "" {
		value = def
	}
	return value
}
