package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	"github.com/rezandry/kejarmimpi/models"
	"github.com/rezandry/kejarmimpi/template"
=======
	"kejarmimpi/models"
	"kejarmimpi/template"
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
	"golang.org/x/crypto/bcrypt"
)

// Login is method post from endUser
//Logic of login
// 1. BE get email and password
// 2. BE check email and password
// 3. BE generate token
// 4. Save token into db
// 5. BE return token and message
func Login(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var user models.User
	//1. get email and password
	c.BindJSON(&user)
	password := user.Password
<<<<<<< HEAD
=======
	// 1.1 becrypt password
	// password := []byte(user.Password)
	// hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	// if err != nil {
	// 	var res models.Response
	// 	res.Code = "401"
	// 	res.Message = "Failed to encrypt password!"
	// 	c.JSON(400, res)
	// }
	// user.Password = string(hashedPassword)
	// log.Println(user.Password)
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
	//2. check email
	if err := db.Select("email, password").Where("email = ?", user.Email).First(&user).Error; err != nil {
		//5. Skip Generate Token and return message
		var res models.Response
		res.Code = "401"
<<<<<<< HEAD
		res.Message = "Email doesn't exist!"
=======
		res.Message = "Your email or password is wrong!"
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
		res.Token = ""
		data := template.Response(&res)
		c.JSON(200, data)
	} else {
		fmt.Println(user.Email)
		fmt.Println(password)
		fmt.Println(user.Password)
		fmt.Println(CheckPasswordHash(password, user.Password))
		if CheckPasswordHash(password, user.Password) {
			//3. Generate Token
			originalText := user.Email + user.Password
			user.Token = GenerateToken(originalText)

			//4. Save token into db
			db.Model(&user).Where("email = ?", user.Email).Update("token", user.Token)
			//5. Return token and message
			var res models.Response
			res.Code = "201"
			res.Message = "Login Success!"
			res.Token = user.Token
			data := template.Response(&res)
			c.JSON(200, data)
			return
		}
		var res models.Response
		res.Code = "402"
<<<<<<< HEAD
		res.Message = "Your Password is wrong!"
=======
		res.Message = "Password salah!"
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
		res.Token = user.Token
		data := template.Response(&res)
		c.JSON(200, data)
	}

}

// CheckPasswordHash is unsing check password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
