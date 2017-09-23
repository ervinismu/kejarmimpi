package models

import (
	"time"
)

//Comment struct for save comment in Post
type Comment struct {
	ID      uint      `form:"id" json:"id"`
	Content string    `form:"content" json:"content"`
	Date    time.Time `form:"date" json:"date"`
	IDPost  uint      `form:"idPost" json:"idPost"`
	IDUser  uint      `form:"idUser" json:"idUser"`
	User    User      `form:"User" json:"User"`
}
