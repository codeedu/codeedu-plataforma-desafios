package utils

import (
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func init() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func ConnectDB(env string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	if env != "test" {
		dsn = os.Getenv("dsn")
		db, err = gorm.Open(os.Getenv("dbType"), dsn)
	} else {
		dsn = os.Getenv("dsnTest")
		db, err = gorm.Open(os.Getenv("dbTypeTest"), dsn)
	}

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("debug") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AutoMigrateDb") == "true" {
		db.AutoMigrate(&domain.User{})
	}

	//defer db.Close()

	return db
}
