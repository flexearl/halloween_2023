package pumpkin

import (
	"log"

	"github.com/flexearl/halloween_2023.git/connections"
)

type PumpkinLink struct {
	DayNumb      int    `json:"daynumber"`
	EmailAddress string `json:"email"`
}

func (pumpkin *PumpkinLink) AddPumpkin() error {
	db := connections.StartDatabase()
	defer db.Close()
	dsn := `INSERT INTO userlinkday (email, daynumber) VALUES (?, ?)`
	_, err := db.Exec(dsn, pumpkin.EmailAddress, pumpkin.DayNumb)
	return err
}

func GetUserPumpkins(email string) []bool {
	db := connections.StartDatabase()
	defer db.Close()
	pumpkins := make([]bool, 4)
	dsn := `SELECT daynumb FROM userlinkday WHERE userlinkday.email=? ORDER BY daynumb ASC`
	rows, err := db.Query(dsn, email)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var dayNumber int
		err = rows.Scan(&dayNumber)
		if err != nil {
			log.Panic(err)
		} else {
			pumpkins[dayNumber-1] = true
		}

	}
	return pumpkins
}
