package database

import (
	"fmt"
	"os"
	"test/models"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() {
	db, err := GetDB()

	if err != nil {
		fmt.Println("No connection to database")
		panic(err)
	} else {
		fmt.Println("Successfully connected!")
		db.AutoMigrate(&models.User{})
	}
}

func GetDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
