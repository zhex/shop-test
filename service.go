package main

func GetAllProducts() ([]Product, error) {
	var products []Product
	result := DB.Find(&products)
	return products, result.Error
}

func CreateProduct(product *Product) error {
	result := DB.Create(product)
	return result.Error
}
