package db

import (
	"LibraryApi/src/libraryService/common"
	"LibraryApi/src/libraryService/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
	"os"
)

func ConnectToDb() *gorm.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:5435/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.Book{}, &model.Category{}, model.Author{}, &common.Role{}, &common.Authorization{})

	return db
}
