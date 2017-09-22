package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rezandry/kejarmimpi/models"
)

// GetPost is method get post from database
func GetPost(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var post []models.Post

	if err := db.Find(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	for i := 0; i < len(post); i++ {
		var user models.User
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
	var user models.User
	var post models.Post
	var res models.Response

	token := c.Request.Header.Get("token")
	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&post)
	post.IDUser = user.ID
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

	var user models.User
	token := c.Request.Header.Get("token")
	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	id := c.Params.ByName("id")
	var post models.Post
	var res models.Response
	db.Where("id = ?", id).First(&post)
	fmt.Println(id)
	fmt.Println(user.ID)
	fmt.Println(post.IDUser)
	if user.ID == post.IDUser {
		if err := db.Delete(&post).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(404)
		} else {
			res.Code = "200"
			res.Message = "Post " + id + " has been deleted."
			c.JSON(200, res)
		}
	} else {
		res.Code = "400"
		res.Message = "You have no access to delete this post"
		c.JSON(400, res)
	}
}

//UpdatePost used for edit post
func UpdatePost(c *gin.Context) {
	var post models.Post
	var res models.Response
	var user models.User
	db := InitDb()
	defer db.Close()

	token := c.Request.Header.Get("token")
	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

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
	} else if post.IDUser == user.ID {
		db.Save(&post)
		res.Code = "200"
		res.Message = "Update post success!"
		c.JSON(200, res)
	} else {
		res.Code = "400"
		res.Message = "You have no access to delete this post"
		c.JSON(200, res)
	}

}
