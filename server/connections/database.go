package connections

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func StartDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:Marley123@tcp(127.0.0.1:3306)/halloween_code")
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Print("Ping  ", err)
	}
	return db
}
