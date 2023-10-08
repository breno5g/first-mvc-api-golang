package main

import (
	"fmt"
	"net/http"

	"github.com/breno5g/mvc-api/routes"
	_ "github.com/lib/pq"
)


func main() {
	routes.LoadRoutes()

	fmt.Println("Server is runing on port 3001")
	http.ListenAndServe(":3001", nil)
}


