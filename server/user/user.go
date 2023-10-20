package user

import (
	"github.com/flexearl/halloween_2023.git/pumpkin"
)

type User struct {
	UserName         string `json:"username"`
	FirstName        string `json:"firstname"`
	Surname          string `json:"surname"`
	PumpkinsAchieved []pumpkin.Pumpkin
	HashedPassword   string
}

func GetUser(userName string) (User, error) {
	//Get user from database
	return User{}, nil
}

func RegisterUser(rawJSON []byte) error {
	//Registers user in the database
	return nil
}
