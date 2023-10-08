package model

import "github.com/breno5g/mvc-api/infra"

type Product struct {
	Id          int
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

func GetProductById(id int) Product {
	db := infra.ConnectWithDatabase()
	defer db.Close()

	productSelect, err := db.Query("SELECT * FROM products WHERE id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	var product Product

	for productSelect.Next() {
		err = productSelect.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			panic(err.Error())
		}
	}

	return product
}

func CreateNewProduct(name string, price float64, description string, quantity int) {
	db := infra.ConnectWithDatabase()
	defer db.Close()

	insertProduct, err := db.Prepare("INSERT INTO products(name, price, description, quantity) VALUES ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(name, price, description, quantity)
}

func DeleteProduct(id int) {
	db := infra.ConnectWithDatabase()
	defer db.Close()

	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
}

func UpdateProduct(id int, name string, price float64, description string, quantity int) {
	db := infra.ConnectWithDatabase()
	defer db.Close()

	updateProduct, err := db.Prepare("UPDATE products SET name=$1, price=$2, description=$3, quantity=$4 WHERE id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, price, description, quantity, id)
}
