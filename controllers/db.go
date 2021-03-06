package controllers

import (
	"fmt"
	"os"

	"github.com/ervinismu/kejarmimpi/models"
	"github.com/jinzhu/gorm"
	//For connect postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//InitDb is for connecting to database
func InitDb() *gorm.DB {
	var err error
	dbhost := os.Getenv("DATABASE_URL")
	if dbhost == "" {
		dbhost = "host=127.0.0.1 user=postgres dbname=kejarmimpidb sslmode=disable password=kejarmimpi"
	}
	var post models.Post
	var user models.User
	var collabs models.Collabs
	var star models.Star
	var comment models.Comment

	db, err := gorm.Open("postgres", dbhost)
	db.AutoMigrate(post, user, collabs, star, &comment)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Success!")
	}
	return db
}
