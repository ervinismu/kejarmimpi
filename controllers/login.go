package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rezandry/kejarmimpi/models"
	"github.com/rezandry/kejarmimpi/template"
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
	c.Bind(&user)
	password := user.Password
	fmt.Println(user.Email)
	fmt.Println(user.Password)
	//2. check email
	if err := db.Select("email, password").Where("email = ?", user.Email).First(&user).Error; err != nil {
		//5. Skip Generate Token and return message
		var res models.Response
		res.Code = "407"
		res.Message = "Email doesn't exist!"
		res.Token = ""
		data := template.Response(&res)
		c.JSON(422, data)
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
		res.Code = "408"
		res.Message = "Your Password is wrong!"
		res.Token = user.Token
		data := template.Response(&res)
		c.JSON(422, data)
	}

}

// CheckPasswordHash is unsing check password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
