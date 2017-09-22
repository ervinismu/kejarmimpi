package controllers

import "github.com/rezandry/kejarmimpi/models"
import "github.com/gin-gonic/gin"
import "fmt"

// CekUser is used for check function debugging
func CekUser(c *gin.Context) {
	var user []models.User
	db := InitDb()
	defer db.Close()
	if err := db.Find(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		data := map[string]interface{}{
			"user": user,
		}
		c.JSON(200, data)
	}
}
