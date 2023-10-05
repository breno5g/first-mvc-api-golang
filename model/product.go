package model

import "github.com/breno5g/mvc-api/infra"

type Product struct {
	Id					int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := infra.ConnectWithDatabase()
	defer db.Close()
	
	productsSelect, err := db.Query("SELECT * FROM products")

	if err != nil {
		panic(err.Error())
	}

	var products []Product

	for productsSelect.Next() {
		var product Product
		err = productsSelect.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}

	return products
}