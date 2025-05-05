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
			{Name: "Apple", Price: 0.99},
			{Name: "Banana", Price: 0.59},
			{Name: "Orange", Price: 1.29},
		}
		db.Create(&products)
	}
}

func init() {
	DB = setupDatabase()
	seedDatabase(DB)
}
