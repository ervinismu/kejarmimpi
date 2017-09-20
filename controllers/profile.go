package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"kejarmimpi/models"
	"kejarmimpi/template"
)

// GetProfile func is for get profile by id
func GetProfile(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user models.User
	// var profile models.Profile
	id := c.Params.ByName("id")

	if err := db.Select("id, name, photo, job, address, status, biograph, education, career").Where("id = ?", id).Find(&user).Error; err != nil {
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
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(200, user)
}
