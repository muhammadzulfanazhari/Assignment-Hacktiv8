package database

import (
	"assignment2/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"postgres",
		"zulfan",
		"hacktiv8",
	)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}

	db.Debug().AutoMigrate(models.Item{}, models.Order{})
	return db
}
