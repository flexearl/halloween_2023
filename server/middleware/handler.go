package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/flexearl/halloween_2023.git/pumpkin"
	"github.com/flexearl/halloween_2023.git/puzzle"
	"github.com/flexearl/halloween_2023.git/user"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

func Register(w http.ResponseWriter, r *http.Request) {
	//Register user in database
	var newUser *user.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		log.Fatal(err)
	}
	newUser.RegisterUser()
}

func AddPumpkin(w http.ResponseWriter, r *http.Request) {
	var newPumpkin pumpkin.PumpkinLink
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newPumpkin)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckUserPumpkins(w http.ResponseWriter, r *http.Request) {
	var emailAddress string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&emailAddress)
	if err != nil {
		log.Fatal(err)
	}
	pumpkins := pumpkin.GetUserPumpkins(emailAddress)
	jsonPumpkins, err := json.Marshal(pumpkins)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonPumpkins)
}

func GetPuzzleContent(w http.ResponseWriter, r *http.Request) {
	var dayNumb int
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dayNumb)
	if err != nil {
		log.Fatal(err)
	}
	content := puzzle.GetPuzzleContent(dayNumb)
	jsonContent, err := json.Marshal(content)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonContent)
}
