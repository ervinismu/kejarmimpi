package controllers

import "golang.org/x/crypto/bcrypt"

<<<<<<< HEAD
//HashPassword is for generate token using AES
=======
//GenerateToken is for generate token using AES
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
