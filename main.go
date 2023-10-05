package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Product struct {
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
	products := []Product{
		{"Milk", "Milk from cow", 2.99, 10},
		{"Eggs", "Eggs from chicken", 1.99, 20},
		{"Bread", "Bread from wheat", 1.49, 5},
		{"Cheese", "Cheese from cow", 0.99, 15},
	}
	templates.ExecuteTemplate(res, "Index", products)
}
