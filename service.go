package main

func GetAllProducts() ([]Product, error) {
	var products []Product
	result := DB.Find(&products)
	return products, result.Error
}
