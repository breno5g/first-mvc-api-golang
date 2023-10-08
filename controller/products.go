package controller

import (
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

	http.Redirect(res, req, "/", 301)
}
