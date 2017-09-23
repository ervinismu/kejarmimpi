package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rezandry/kejarmimpi/models"
)

//Star is used for get content
func Star(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var star []models.Star

	if err := db.Find(&star).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		data := map[string]interface{}{
			"star": star,
		}
		c.JSON(200, data)
	}
}

//CreateStar is used to input content star
func CreateStar(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var star models.Star
	var res models.Response

	c.Bind(&star)

	if err := db.Create(&star).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		res.Code = "200"
		res.Message = "Post Success!"
		c.JSON(200, res)
	}
}
