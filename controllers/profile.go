package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rezandry/kejarmimpi/models"
	"github.com/rezandry/kejarmimpi/template"
)

// GetProfile func is for get profile by id
func GetProfile(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user models.User
	// var profile models.Profile
	token := c.Request.Header.Get("token")

	if err := db.Select("id, name, photo, job, address, status, biograph, education, career").Where("token = ?", token).Find(&user).Error; err != nil {
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
	// id := c.Params.ByName("id")
	token := c.Request.Header.Get("token")
	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.Bind(&user)
	db.Save(&user)
	user.Password = ""
	user.Token = ""
	c.JSON(200, user)
}
