package controller

import (
	"html/template"
	"net/http"

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
