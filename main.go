package main

import (
	"net/http"

	"github.com/pequenojoohn/routes"
)

func main() {
	routes.LoadingRoutes()
	http.ListenAndServe(":8000", nil)
}
