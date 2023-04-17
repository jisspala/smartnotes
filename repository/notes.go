package repository

import (
	"fmt"
	"os"
	"smartnotes/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// This variable will be used to store the connection to the PostgreSQL
var db *gorm.DB

// initializes a connection to a PostgreSQL and do auto migrates a Note model.
func InitDb() {
	var err error
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	host := os.Getenv("POSTGRES_HOST")
	dbname := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, user, password, dbname, port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.Note{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}

// The function returns a pointer to a GORM database instance.
func GetDB() *gorm.DB {
	return db
}
