package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"kejarmimpi/models"
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
	c.BindJSON(&post)
	post.Date = time.Now()
	if err := db.Create(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, post)
	}
}

//DeletePost used for delete post that exist before
func DeletePost(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var post models.Post
	if err := db.Where("id = ?", id).Delete(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(404)
	} else {
		c.JSON(200, gin.H{"post #" + id: "deleted"})
	}

}

//UpdatePost used for edit post
func UpdatePost(c *gin.Context) {
	var post models.Post
	db := InitDb()
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&post)
	db.Save(&post)
	c.JSON(200, post)
}
