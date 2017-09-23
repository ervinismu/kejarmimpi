package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ervinismu/kejarmimpi/models"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {

	// define db
	db := InitDb()
	defer db.Close()

	var user models.User
	var comment models.Comment
	var res models.Response

	// get data when send
	c.Bind(&comment)

	// get token and id post
	postID := c.Params.ByName("id") // type data string dibuat uint
	token := c.Request.Header.Get("token")

	// convert string to uint in ID post
	u64, err := strconv.ParseUint(postID, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	x := uint(u64)

	// nggoleki id user
	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	// Check if user send blank content
	if comment.Content == "" {
		res.Code = "400"
		res.Message = "Content must not blank!"
		c.JSON(400, res)
	} else {
		t := time.Now().Add(7 * time.Hour)
		comment.Date = t

		// get IDUser and IDPost
		// from comment
		comment.IDUser = user.ID
		comment.IDPost = x

		if err := db.Create(&comment).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			res.Code = "200"
			res.Message = "Comment Succes!"
			c.JSON(200, res)
		}

	}

}
