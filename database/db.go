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
		fmt.Println("No connection to database", err)
		panic(err)
	} else {
		fmt.Println("Successfully connected!")
		db.AutoMigrate(&models.User{})
	}
}

func GetDB() (*gorm.DB, error) {
	// host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=pgadmin dbname=postgres password=postgres sslmode=disable")

	db = db.Exec("CREATE DATABASE test_db;")

	psqlInfo := fmt.Sprintf("port=%s user=%s "+"password=%s dbname=%s sslmode=disable", port, user, password, dbname)
	DB, error := gorm.Open("postgres", psqlInfo)

	defer db.Close()

	if error != nil {
		return nil, err
	}
	return DB, nil
}
