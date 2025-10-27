package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/flexearl/halloween_2023.git/pumpkin"
	"github.com/flexearl/halloween_2023.git/puzzle"
	"github.com/flexearl/halloween_2023.git/user"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

func GetUserPuzzleInput(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting user puzzle input")
	userIDStr := r.URL.Query().Get("userid")
	dayNumberStr := r.URL.Query().Get("daynumber")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}
	dayNumb, err := strconv.Atoi(dayNumberStr)
	if err != nil {
		http.Error(w, "Invalid Day", http.StatusBadRequest)
		return
	}

	puzzleInput := puzzle.GetUserPuzzleInput(dayNumb, userID)
	response := map[string]interface{}{
		"message":      "Successful Request",
		"puzzle_input": puzzleInput,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
	fmt.Println("Finished")
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Registering User")
	//Register user in database
	var newUser *user.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		log.Fatal(err)
	}
	userID := newUser.RegisterUser()
	puzzle.AddUserPuzzleInput(userID)
	response := map[string]interface{}{
		"message": "User created successfully",
		"userID":  userID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
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
	fmt.Println("Getting Puzzle Content")
	var dayNumb int

	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/get_puzzle_content/")
	dayNumb, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid puzzle number", http.StatusBadRequest)
		return
	}

	content := puzzle.GetPuzzleContent(dayNumb)
	jsonContent, err := json.Marshal(content)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonContent)
}
