package models

import (
	"time"
)

//Post is struct for save data post
type Post struct {
	ID         uint      `json:"id"`
	Content    string    `json:"content"`
	Date       time.Time `json:"date"`
	Photo      string    `json:"photo"`
	IDUser     uint      `json:"idUser"`
	User       User      `json:"User"`
	IDCategory uint      `json:"idCategory"`
}
