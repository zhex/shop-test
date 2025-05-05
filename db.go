package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func setupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})
	return db
}

func seedDatabase(db *gorm.DB) {
	var count int64
	db.Model(&Product{}).Count(&count)
	if count == 0 {
		products := []Product{
			{ID: 1, Name: "Chrome Toaster", Price: 100},
			{ID: 2, Name: "Copper Kettle", Price: 49.99},
			{ID: 3, Name: "Mixing Bowl", Price: 20},
		}
		db.Create(&products)
	}
}

func init() {
	DB = setupDatabase()
	seedDatabase(DB)
}
