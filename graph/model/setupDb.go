package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB database object
var DB *gorm.DB

// InitDb initializes a db
func InitDb() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(&Item{}, &User{}, &Order{})

	DB = database
}
