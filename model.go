package main

type Product struct {
	ID    int     `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0"`
}
