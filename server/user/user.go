package user

import (
	"github.com/flexearl/halloween_2023.git/connections"
	"github.com/flexearl/halloween_2023.git/password"
)

type User struct {
	UserName         string `json:"username"`
	FirstName        string `json:"firstname"`
	Surname          string `json:"surname"`
	EmailAddress     string `json:"email"`
	PumpkinsAchieved []bool
	HashedPassword   string `json:"password"`
}

func GetUser(userName string) (User, error) {
	//Get user from database
	newUser := User{}
	dsn := `SELECT username, email, hashedpassword, firstname, surname 
	FROM user WHERE user.username = ?`
	db := connections.StartDatabase()
	row := db.QueryRow(dsn, userName)
	err := row.Scan(&newUser.UserName, &newUser.EmailAddress,
		&newUser.HashedPassword, &newUser.FirstName, &newUser.Surname)
	return newUser, err
}

func (u *User) RegisterUser() error {
	//Registers user in the database
	u.HashedPassword = password.HashPassword(u.HashedPassword)
	dsn := `INSERT INTO TABLE user (username, email, hashedpassword, firstname, surname)
	VALUES (?, ?, ?, ?, ?)`
	db := connections.StartDatabase()
	_, err := db.Exec(dsn, u.UserName, u.EmailAddress,
		u.HashedPassword, u.FirstName, u.Surname)
	return err
}
