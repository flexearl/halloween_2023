package routes

import (
	"net/http"
)

func HandleRouting() {
	http.Handle("/", Home)
}
