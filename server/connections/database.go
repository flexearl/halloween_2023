package connections

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func StartDatabase() *sql.DB {
	db, err := sql.Open("mysql", "flex:Marley123@tcp(halloween-2023.c5zvii8avkbe.eu-west-2.rds.amazonaws.com:3306)/initial_db")
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Print("Ping  ", err)
	}
	return db
}
