package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"go-web-app.com/routes"
)



func main() {
	routes.InitializeRoutes()
	http.ListenAndServe(":8000", nil)
}

