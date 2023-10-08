package routes

import (
	"net/http"

	"github.com/breno5g/mvc-api/controller"
)

func LoadRoutes() {
	http.HandleFunc("/", controller.IndexHandler)

}
