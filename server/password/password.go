package password 

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error){
	bytes, err := GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

