package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/breno5g/mvc-api/model"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	products := model.GetAllProducts()

	templates.ExecuteTemplate(res, "Index", products)
}

func NewProductHandler(res http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(res, "New-Product", nil)
}

func CreateProductHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		name := req.FormValue("name")
		price := req.FormValue("price")
		description := req.FormValue("description")
		quantity := req.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			panic(err.Error())
		}

		convertedQuantity, err := strconv.Atoi(quantity)

		if err != nil {
			panic(err.Error())
		}

		model.CreateNewProduct(name, convertedPrice, description, convertedQuantity)
	}

	http.Redirect(res, req, "/", http.StatusMovedPermanently)
}

func DeleteProductHandler(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")

	convertedId, err := strconv.Atoi(id)

	if err != nil {
		panic(err.Error())
	}

	model.DeleteProduct(convertedId)

	http.Redirect(res, req, "/", http.StatusMovedPermanently)
}

func EditProductHandler(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")

	convertedId, err := strconv.Atoi(id)

	if err != nil {
		panic(err.Error())
	}

	product := model.GetProductById(convertedId)

	templates.ExecuteTemplate(res, "Edit-Product", product)
}

func UpdateProductHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		fmt.Println("UPDATE")

		id := req.FormValue("id")
		name := req.FormValue("name")
		price := req.FormValue("price")
		description := req.FormValue("description")
		quantity := req.FormValue("quantity")

		convertedId, err := strconv.Atoi(id)

		if err != nil {
			panic(err.Error())
		}

		convertedPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			panic(err.Error())
		}

		convertedQuantity, err := strconv.Atoi(quantity)

		if err != nil {
			panic(err.Error())
		}

		model.UpdateProduct(convertedId, name, convertedPrice, description, convertedQuantity)
	}

	http.Redirect(res, req, "/", http.StatusMovedPermanently)
}
