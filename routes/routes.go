package routes

import (
	"net/http"

	"github.com/breno5g/mvc-api/controller"
)

func LoadRoutes() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", controller.IndexHandler)
	http.HandleFunc("/new-product", controller.NewProductHandler)
	http.HandleFunc("/create-product", controller.CreateProductHandler)
}
