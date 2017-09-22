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
)

// GetProfile func is for get profile by id
func GetProfile(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user models.User
	// var profile models.Profile
<<<<<<< HEAD
	token := c.Request.Header.Get("token")

	if err := db.Select("id, name, photo, job, address, status, biograph, education, career").Where("token = ?", token).Find(&user).Error; err != nil {
=======
	id := c.Params.ByName("id")

	if err := db.Select("id, name, photo, job, address, status, biograph, education, career").Where("id = ?", id).Find(&user).Error; err != nil {
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		data := template.Profile(&user)
		c.JSON(200, data)
	}
}

//UpdateProfile is func for edit user profile
func UpdateProfile(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user models.User
<<<<<<< HEAD
	// id := c.Params.ByName("id")
	token := c.Request.Header.Get("token")
	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
=======
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&user)
	db.Save(&user)
<<<<<<< HEAD
	user.Password = ""
	user.Token = ""
=======
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
	c.JSON(200, user)
}
