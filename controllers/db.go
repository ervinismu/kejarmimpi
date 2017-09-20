package controllers

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"kejarmimpi/models"
	//For connect postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//InitDb is for connecting to database
func InitDb() *gorm.DB {
	// var db *gorm.DB
	dbhost := os.Getenv("DATABASE_URL")
	fmt.Println(dbhost)
	//if you are local, use 
	//db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=namadatabase sslmode=disable password=password")
	db, err := gorm.Open("postgres", dbhost)
	db.AutoMigrate(models.User{}, models.Post{})
	if err != nil {
		fmt.Println("Ga asek")
		fmt.Println(err)
	} else {
		fmt.Println("Connection Success!")
	}
	return db
}
