package routes

import (
	"net/http"

	"github.com/flexearl/halloween_2023.git/middleware"
)

func HandleRouting() {
	http.HandleFunc("/", middleware.Home)
	http.HandleFunc("/register_user", middleware.Register)
	http.HandleFunc("/check_user_pumpkins", middleware.CheckUserPumpkins)
	http.HandleFunc("/add_pumpkin", middleware.AddPumpkin)
}
