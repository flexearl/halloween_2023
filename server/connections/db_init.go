package connections

import "fmt"

func init_User() {
	dsn := `CREATE TABLE user (
		username varchar(50),
		email varchar(50),
		hashedpassword varchar(50),
		firstname varchar(50),
		surname varchar(50)
	)`
	fmt.Println(dsn)
}

func init_puzzle() {
	dsn := `CREATE TABLE puzzle (
		number int,
		content varchar(255)
	)`
	fmt.Println(dsn)
}

func init_userlinkday() {
	dsn := `CREATE TABLE userlinkday (
		email varchar(50),
		daynumber int
	)`
	fmt.Println(dsn)
}
