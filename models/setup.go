package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	if os.Getenv("APP_ENV") != "production" {
		dotEnvErr := godotenv.Load()
		if dotEnvErr != nil {
			log.Fatal("Error loading .env file")
		}
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PWD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=prefer TimeZone=Asia/Jakarta", host, user, pwd, dbname, dbport)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection error!")
	}

	db.AutoMigrate(&Todo{})

	fmt.Println("DB connection success!")

	DB = db
}