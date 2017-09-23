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

	if err := db.Raw("SELECT * FROM posts ORDER BY id DESC").Scan(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	for i := 0; i < len(post); i++ {
		var user models.User
		if err := db.Where("id = ?", post[i].IDUser).Find(&user).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		post[i].CollabsCount = CountCollabs(post[i].ID)
		fmt.Println(post[i].CollabsCount)
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
	// fmt.Println(c.PostForm("stuff"))
	// fmt.Println(c.PostForm("astuff"))
	// c.JSON(http.StatusOK, gin.H{"response": c.PostForm("stuff")})

	// type Req struct {
	// 	Content string `form:"content" json:"content" binding:"required"`
	// 	Photo   string `form:"photo" json:"photo" binding:"required"`
	// }
	// var req Req

	db := InitDb()
	defer db.Close()
	var user models.User
	var post models.Post
	c.Bind(&post)
	var res models.Response

	// type PostReq struct {
	// 	Content string `form:"content" json:"content" binding:"required"`
	// 	Photo   string `form:"photo" json:"photo" binding:"required"`
	// }

	token := c.Request.Header.Get("token")
	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	// c.BindJSON(&post)
	// var form PostReq
	c.Bind(&post)
	fmt.Println(post.Content)

	if post.Content == "" {
		fmt.Println("Ga entuk")
		fmt.Println("Ga entuk")
	}

	if post.Content == "" && post.Photo == "" {
		res.Code = "401"
		res.Message = "Content must not blank!"
		c.JSON(400, res)
	} else {
		t := time.Now().Add(7 * time.Hour)
		post.Date = t
		post.IDUser = user.ID
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
