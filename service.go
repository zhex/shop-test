package main

import (
	"gorm.io/gorm"
)

func GetAllProducts() ([]Product, error) {
	var products []Product
	result := DB.Find(&products)
	return products, result.Error
}

func CreateProduct(product *Product) error {
	result := DB.Create(product)
	return result.Error
}

func GetProductByID(id int) (*Product, error) {
	var product Product
	result := DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func GetProductByIDs(ids []int) (map[int]Product, error) {
	var products []Product
	result := DB.Where("id IN ?", ids).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	productMap := make(map[int]Product)
	for _, p := range products {
		productMap[p.ID] = p
	}
	return productMap, nil
}

func CreateSale(req SaleRequest) (*Sale, error) {
	var sale Sale
	err := DB.Transaction(func(tx *gorm.DB) error {
		var ids []int
		for _, item := range req.LineItems {
			ids = append(ids, item.ProductID)
		}
		productMap, err := GetProductByIDs(ids)
		if err != nil {
			return err
		}

		var lineItems []LineItem
		total := 0.0
		for _, item := range req.LineItems {
			product, ok := productMap[item.ProductID]
			if !ok {
				return gorm.ErrRecordNotFound
			}
			lineTotal := product.Price * float64(item.Quantity)
			lineItem := LineItem{
				SaleID:    sale.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     product.Price,
				Total:     lineTotal,
			}
			lineItems = append(lineItems, lineItem)
			total += lineTotal
		}

		discount := req.Discount
		if discount > total {
			// Cap discount to total
			discount = total
		}
		var distributed float64
		for i := range lineItems {
			if i == len(lineItems)-1 {
				// Last item gets the remainder to avoid floating point issues
				lineItems[i].Discount = discount - distributed
			} else {
				prop := lineItems[i].Total / total
				lineItems[i].Discount = roundToTwo(prop * discount)
				distributed += lineItems[i].Discount
			}
			lineItems[i].Total = lineItems[i].Total - lineItems[i].Discount
		}

		sale.LineItems = lineItems
		sale.Total = total - discount
		sale.Discount = discount

		return tx.Create(&sale).Error
	})
	if err != nil {
		return nil, err
	}
	return &sale, nil
}

func roundToTwo(val float64) float64 {
	return float64(int(val*100+0.5)) / 100
}
