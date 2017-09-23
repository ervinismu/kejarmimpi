package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rezandry/kejarmimpi/models"
)

//Collabs is for collabs
func Collabs(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var collabs models.Collabs
	var user models.User
	var post models.Post
	var res models.Response

	idPost := c.Params.ByName("id")
	token := c.Request.Header.Get("token")
	fmt.Println(token)
	if err := db.Select("token, id").Where("token = ?", token).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	fmt.Println(user.ID)

	if err := db.Select("id, id_user").Where("id = ?", idPost).First(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	if err := db.Select("id").Where("id_post = ? AND id_user = ?", idPost, user.ID).First(&collabs).Error; err != nil {
		fmt.Println("create")
		collabs.IDUPost = post.IDUser
		collabs.IDPost = post.ID
		collabs.IDUser = user.ID
		res.Code = "200"
		res.Message = "Collabs"
		res.Status = true
		db.Save(&collabs)
		c.JSON(200, res)
	} else {
		fmt.Println("delete")
		db.Delete(&collabs)
		res.Code = "200"
		res.Message = "Uncollabs"
		res.Status = false
		c.JSON(200, res)
	}

}

func NotifCollabs(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user models.User
	var collabs []models.Collabs
	token := c.Request.Header.Get("token")
	fmt.Println(token)
	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	if err := db.Where("id_u_post = ?", user.ID).Find(&collabs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else if len(collabs) > 0 {
		for i := 0; i <= len(collabs); i++ {
			var userc models.User
			if err := db.Where("id = ?", collabs[i].IDUser).First(&userc).Error; err != nil {
				c.AbortWithStatus(404)
				fmt.Println(err)
			}
			collabs[i].Message = userc.Name + "Interest to collabs with you."
		}
		data := map[string]interface{}{
			"notif": collabs,
		}
		c.JSON(200, data)
		c.Abort()
		c.JSON(200, "You have notification")
		return
	}
	fmt.Println(len(collabs))
	c.JSON(400, "You have no notification")

}
