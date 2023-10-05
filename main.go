package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func connectWithDatabase() *sql.DB {
	connStr := "user=admin password=root dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error() )
	}
	return db
}

type Product struct {
	Id					int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", indexHandler)

	fmt.Println("Server is runing on port 3001")
	http.ListenAndServe(":3001", nil)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	db := connectWithDatabase()
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

	fmt.Println(products)

	templates.ExecuteTemplate(res, "Index", products)
}
