package controllers

import "golang.org/x/crypto/bcrypt"

//HashPassword is for generate token using AES
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
