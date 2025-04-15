package database;

import (
	"github.com/leonardoavalerio/api-files/internal/model"

	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
)

func initDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=2319 dbname=sys_expenses port=5432 sslmode=disable";

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	db.AutoMigrate(&model.File{});

	return db;
}

var Db = initDatabase();