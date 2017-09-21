package controllers

import (
	"regexp"

	"github.com/ervinismu/kejarmimpi/models"
	"github.com/ervinismu/kejarmimpi/template"
	"github.com/gin-gonic/gin"
)

// var db *gorm.DB

// Register func for register new user
func Register(c *gin.Context) {
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	db := InitDb()
	defer db.Close()
	var user models.User
	var res models.Response
	// user := models.User{}
	// get data
	c.BindJSON(&user)
	// Check
	// if user.Email == string("") {
	// 	response.Code = "401"
	// 	response.Message = "Email can not be empty"
	// 	c.JSON(400, response)
	// 	return
	// } else if user.Name == string("") {
	// 	response.Code = "401"
	// 	response.Message = "Name can not be empty"
	// 	c.JSON(400, response)
	// 	return
	// } else if user.Password+user.PasswordConfirm == string("") {
	// 	response.Code = "401"
	// 	response.Message = "Password can not be empty"
	// 	c.JSON(400, response)
	// 	return
	// }
	// Check Valid email
	if user.Name == string("") {
		res.Code = "401"
		res.Message = "can not be empty"
		c.JSON(400, res)
		// } else if user.Name == string(""){
		// 	response.Code = "401"
		// 	response.Message = "can not be empty"
		// 	c.JSON(400, response)
		// } else if user.Password + user.PasswordConfirm == string(""){
		// 	response.Code = "401"
		// 	response.Message = "can not be empty"
		// 	c.JSON(400, response)
	} else if user.Password == string("") {
		res.Code = "403"
		res.Message = "password can not be empty"
		c.JSON(400, res)
	} else if !emailRegexp.MatchString(user.Email) {
		res.Code = "401"
		res.Message = "invalid format email!"
		c.JSON(400, res)
	} else if user.Password != user.PasswordConfirm {
		res.Code = "401"
		res.Message = "Password not same!"
		c.JSON(400, res)
	} else if err := db.Where("email = ?", user.Email).First(&user).Error; err != nil {
		plainPassword := user.Password
		hashedPassword, err := HashPassword(user.Password)
		if err != nil {
			res.Code = "401"
			res.Message = "Failed to encrypt password!"
			c.JSON(400, res)
		}
		user.Password = string(hashedPassword)
		user.PasswordConfirm = string("")
		// Insert into database after encrypt
		db.Save(&user)
		if CheckPasswordHash(plainPassword, user.Password) {
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

	} else {
		res.Code = "401"
		res.Message = "Email has been used!"
		c.JSON(400, res)

	}
}
