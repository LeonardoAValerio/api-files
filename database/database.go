package database;

import (
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
	"github.com/leonardoavalerio/api-files/model"
	"os"
	"fmt"
)

func initDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	);
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	db.AutoMigrate(&model.File{});

	return db;
}

var Db = initDatabase();