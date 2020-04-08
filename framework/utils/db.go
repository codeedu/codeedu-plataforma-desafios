package utils

import (
	"fmt"
	"github.com/codeedu/codeedu/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func ConnectDB() *gorm.DB {

	//Load environmenatal variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("dsn")

	//Define DB connection string
	dbURI := fmt.Sprintf(dsn)

	//connect to db URI
	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	// close db when not in use
	//defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&domain.User{})

	fmt.Println("Successfully connected!")
	return db
}
