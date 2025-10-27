package puzzle

import (
	"fmt"
	"log"

	"github.com/flexearl/halloween_2023.git/connections"
)

func GetPuzzleContent(dayNumb int) string {
	dsn := `SELECT content FROM puzzle WHERE puzzle.daynumber = ?`
	db := connections.StartDatabase()
	defer db.Close()
	row := db.QueryRow(dsn, dayNumb)
	var content string
	err := row.Scan(&content)
	if err != nil {
		log.Print("Could not scan content", err)
	}
	return content
}

func GetUserPuzzleInput(dayNumb, userId int) string {
	dsn := `SELECT puzzle_input FROM user_puzzle_input WHERE daynumber = ? AND user_id = ?`
	db := connections.StartDatabase()
	defer db.Close()
	row := db.QueryRow(dsn, dayNumb, userId)
	var input string
	err := row.Scan(&input)
	if err != nil {
		log.Print("Could not scan Puzzle Input", err)
	}
	fmt.Println("Input", input)
	return input
}

func AddUserPuzzleInput(userId int) {
	randomSelectionDSN := `SELECT puzzle_input FROM puzzle_input_answer WHERE daynumber = ? ORDER BY RAND() LIMIT 1;`
	addToUserPuzzleInputDSN := `INSERT INTO user_puzzle_input (daynumber, puzzle_input, user_id) VALUES (?, ? , ?)`
	db := connections.StartDatabase()
	defer db.Close()
	for i := 1; i <= 3; i++ {
		row := db.QueryRow(randomSelectionDSN, i)
		var input string
		err := row.Scan(&input)
		if err != nil {
			log.Print(err)
		}
		fmt.Println("Input:", input)
		_, err = db.Exec(addToUserPuzzleInputDSN, i, input, userId)
		if err != nil {
			log.Print(err)
		}

	}
}
