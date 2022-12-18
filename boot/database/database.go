package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB
var err error

func init() {

	err := godotenv.Load()

	if err != nil {
		log.Printf("Erro load env \n details: %s")
		return
	}
}

func newConnection() *gorm.DB {

	if db != nil {
		return db
	}

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		),
	}))

	if err != nil {
		log.Fatalf("error when trying to connect bank \n details: %s", err)
	}

	return db
}

func GetConnection() *gorm.DB {

	return newConnection()
}
