package app

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ptdrpg/wallet/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connexion() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file", err)
	}
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPass, dbName, dbPort)
	db, errors := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errors != nil {
		log.Fatalf("failed to connect to the database: %v", errors)
	}

	db.AutoMigrate(
		&entity.Wallet{},
		&entity.Transaction{},
	)

	DB = db
}