package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func SetUp() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("failed load file env")
	}

	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbport, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed konek ke database")
	}
	return db
}

func CloseDatabase(db *gorm.DB) {
	dbsql, err := db.DB()
	if err != nil {
		panic("failed to close connection")
	}
	dbsql.Close()
}
