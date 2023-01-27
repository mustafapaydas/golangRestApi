package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"libraryApi/src/libraryService/model"
	"log"
)

func ConnectToDb() *gorm.DB {
	dbURL := "postgres://library:library@localhost:5432/library"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.Book{})

	return db
}
