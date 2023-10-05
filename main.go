package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/breno5g/mvc-api/model"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", indexHandler)

	fmt.Println("Server is runing on port 3001")
	http.ListenAndServe(":3001", nil)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	products := model.GetAllProducts()

	templates.ExecuteTemplate(res, "Index", products)
}
