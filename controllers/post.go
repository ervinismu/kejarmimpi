package controllers

import (
	"fmt"
	"time"

	"github.com/ervinismu/kejarmimpi/models"
	"github.com/gin-gonic/gin"
)

// GetPost is method get post from database
func GetPost(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var post []models.Post
	var user models.User

	if err := db.Find(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	for i := 0; i < len(post); i++ {
		if err := db.Where("id = ?", post[i].IDUser).Find(&user).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		user.Password = ""
		user.Token = ""
		user.PasswordConfirm = ""
		post[i].User = user
	}
	data := map[string]interface{}{
		"post": post,
	}
	c.JSON(200, data)
}

//CreatePost used for create a new post
func CreatePost(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var post models.Post
	var res models.Response
	c.BindJSON(&post)
	if post.Content == "" && post.Photo == "" {
		res.Code = "401"
		res.Message = "Content must not blank!"
		c.JSON(400, res)
	} else {
		post.Date = time.Now()
		if err := db.Create(&post).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			res.Code = "200"
			res.Message = "Post Success!"
			c.JSON(200, res)
		}
	}
}

//DeletePost used for delete post that exist before
func DeletePost(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var post models.Post
	var res models.Response
	if err := db.Where("id = ?", id).Delete(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(404)
	} else {
		res.Code = "200"
		res.Message = "Post " + id + " has been deleted."
		c.JSON(200, res)
	}

}

//UpdatePost used for edit post
func UpdatePost(c *gin.Context) {
	var post models.Post
	var res models.Response
	db := InitDb()
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&post)

	if post.Content == "" && post.Photo == "" {
		res.Code = "401"
		res.Message = "Content must not blank!"
		c.JSON(400, res)
		c.AbortWithStatus(400)
	} else {
		db.Save(&post)
		res.Code = "200"
		res.Message = "Update post success!"
		c.JSON(200, res)
	}

}
