package model

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}

func (user *User) Create() {
	hashedPassword, _ := HashPassword(user.Password)
	user.Password = string(hashedPassword)

	result := GetDB().Create(user)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserIdByUsername(username string) (string, error) {
	user := User{Username: username}
	err := GetDB().First(&user).Error
	if err != nil {
		log.Fatal(err)
	}
	return user.ID, nil
}

func (user *User) Authenticate() bool {
	usr := User{Username: user.Username}
	err := GetDB().First(&usr).Error

	if err != nil {
		log.Fatal(err)
		return false
	}

	return CheckPasswordHash(user.Password, usr.Password)
}
