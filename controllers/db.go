package controllers

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/rezandry/kejarmimpi/models"
	//For connect postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//InitDb is for connecting to database
func InitDb() *gorm.DB {
	var err error
	dbhost := os.Getenv("DATABASE_URL")
	if dbhost == "" {
		dbhost = "host=127.0.0.1 user=postgres dbname=kejarmimpi sslmode=disable password=postgresreza"
	}
	var post models.Post
	var user models.User
	var collabs models.Collabs
	var star models.Star

	db, err := gorm.Open("postgres", dbhost)
	db.AutoMigrate(post, user, collabs, star)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Success!")
	}
	return db
}
