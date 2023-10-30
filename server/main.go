package main

import (
	"log"
	"net/http"

	"github.com/flexearl/halloween_2023.git/connections"
	"github.com/flexearl/halloween_2023.git/routes"
)

func main() {
	//svr := http.NewServeMux()
	routes.HandleRouting()
	connections.StartDatabase()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
