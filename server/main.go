package main

import (
	"github.com/flexearl/halloween_2023.git/routes"
	"log"
	"net/http"
)

func main() {
	svr := http.NewServeMux()
	routes.HandleRouting()
	err := http.ListenAndServe(":8000", svr)
	if err != nil {
		log.Fatal(err)
	}
}
