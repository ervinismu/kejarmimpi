package controllers

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
<<<<<<< HEAD
	"github.com/rezandry/kejarmimpi/models"
=======
	"kejarmimpi/models"
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
	//For connect postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//InitDb is for connecting to database
func InitDb() *gorm.DB {
<<<<<<< HEAD
	var err error
	dbhost := os.Getenv("DATABASE_URL")
	if dbhost == "" {
		dbhost = "host=127.0.0.1 user=postgres dbname=kejarmimpi sslmode=disable password=postgresreza"
	}
	var post models.Post
	var user models.User
	db, err := gorm.Open("postgres", dbhost)
	db.AutoMigrate(post, user)
	if err != nil {
=======
	// var db *gorm.DB
	dbhost := os.Getenv("DATABASE_URL")
	fmt.Println(dbhost)
	//if you are local, use 
	//db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=namadatabase sslmode=disable password=password")
	db, err := gorm.Open("postgres", dbhost)
	db.AutoMigrate(models.User{}, models.Post{})
	if err != nil {
		fmt.Println("Ga asek")
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
		fmt.Println(err)
	} else {
		fmt.Println("Connection Success!")
	}
	return db
}
