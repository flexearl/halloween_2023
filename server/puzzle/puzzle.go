package puzzle

import (
	"github.com/flexearl/halloween_2023.git/connections"
)

func GetPuzzleContent(dayNumb int) string {
	dsn := `SELECT content FROM puzzle WHERE puzzle.daynumber = ?`
	db := connections.StartDatabase()
	defer db.Close()
	row := db.QueryRow(dsn, dayNumb)
	var content string
	row.Scan(&content)
	return content
}
