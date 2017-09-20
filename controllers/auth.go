package controllers

import (
	"regexp"

	"kejarmimpi/models"
	// "github.com/jinzhu/gorm"
	//jwt "gopkg.in/appleboy/gin-jwt.v2"
	"github.com/gin-gonic/gin"
	// "github.com/lib/pq"
)

// var db *gorm.DB

// Response Struct for error response
type Response struct {
	Code    string `json:"error_code"`
	Message string `json:"message"`
}

// Register func for register new user
func Register(c *gin.Context) {
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	response := Response{}
	db := InitDb()
	defer db.Close()
	var user models.User
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
		response.Code = "401"
		response.Message = "can not be empty"
		c.JSON(400, response)
		// } else if user.Name == string(""){
		// 	response.Code = "401"
		// 	response.Message = "can not be empty"
		// 	c.JSON(400, response)
		// } else if user.Password + user.PasswordConfirm == string(""){
		// 	response.Code = "401"
		// 	response.Message = "can not be empty"
		// 	c.JSON(400, response)
	} else if user.Password == string("") {
		response.Code = "403"
		response.Message = "password can not be empty"
		c.JSON(400, response)
	} else if !emailRegexp.MatchString(user.Email) {
		response.Code = "401"
		response.Message = "invalid format email!"
		c.JSON(400, response)
	} else if user.Password != user.PasswordConfirm {
		response.Code = "401"
		response.Message = "Password not same!"
		c.JSON(400, response)
	} else if err := db.Where("email = ?", user.Email).First(&user).Error; err != nil {
		hashedPassword, err := HashPassword(user.Password)
		if err != nil {
			response.Code = "401"
			response.Message = "Failed to encrypt password!"
			c.JSON(400, response)
		}
		user.Password = string(hashedPassword)
		// user.PasswordConfirm = string("")
		// Insert into database after encrypt
		db.Save(&user)
		// Response
		response.Code = "200"
		response.Message = "Register Succes!"
		c.JSON(200, response)

	} else {
		response.Code = "401"
		response.Message = "Email has been used!"
		c.JSON(400, response)

	}
}
