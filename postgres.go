package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

func DatabaseConncetion() {
	host := "localhost"
	port := "5432"
	username := "postgres"
	password := "root"
	database := "postgres"

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, database, password)
	fmt.Println("dsn ---> ", dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(Book{})

	if err != nil {
		log.Fatal("Error connecting to database... ", err)
	}
	fmt.Println("Database connection successful...")
}
