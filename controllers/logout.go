package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rezandry/kejarmimpi/models"
	"github.com/rezandry/kejarmimpi/template"
)

//Logout is func for logout from accout
//Logic of logout
//1. BE get token
//2. Delete token
//3. BE return message
func Logout(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var user models.User
	var res models.Response
	//1. BE get token
	token := c.Request.Header.Get("token")

	if token == "" {
		res.Code = "401"
		res.Message = "You are not login!"
		data := template.Response(&res)
		c.JSON(200, data)
	} else if err := db.Select("token").Where("token = ?", token).First(&user).Error; err != nil {
		//3. BE return message
		res.Code = "402"
		res.Message = "Your token invalid"
		data := template.Response(&res)
		c.JSON(200, data)
	} else {
		db.Model(&user).Select("token").Update(map[string]interface{}{"token": ""})
		res.Code = "200"
		res.Message = "Logout Success!"
		data := template.Response(&res)
		c.JSON(200, data)
	}
}
