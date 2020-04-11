package utils

import (
	"fmt"
	"log"
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func ConnectDB() *gorm.DB {

	//Load environmenatal variable
	dsn, err := Getenv("dsn")

	if err != nil {
		log.Fatal("Error loading environment variable")
	}

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

	Printf("Successfully connected!")
	return db
}
