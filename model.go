package main

type Product struct {
	ID    int     `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0"`
}

type Sale struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	LineItems []LineItem `json:"lineItems" gorm:"foreignKey:SaleID"`
	Total     float64    `json:"total"`
	Discount  float64    `json:"discount"`
}

type LineItem struct {
	ID        int     `json:"id" gorm:"primaryKey"`
	SaleID    int     `json:"-"`
	ProductID int     `json:"productId"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Total     float64 `json:"total"`
	Discount  float64 `json:"discount"`
}

type SaleLineItem struct {
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type SaleRequest struct {
	LineItems []SaleLineItem `json:"lineItems"`
	Discount  float64        `json:"discount"`
}
