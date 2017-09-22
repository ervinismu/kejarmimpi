package main

import (
	"os"

	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	"github.com/rezandry/kejarmimpi/controllers"
	"github.com/rezandry/kejarmimpi/middleware"
=======
	"kejarmimpi/controllers"
	"kejarmimpi/middleware"
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
)

func main() {

	// This function is for AutoMigrate
	// db := configs.InitDb()
	// defer db.Close()

	// var post models.Post
	// var user models.User
	// db.AutoMigrate(&user)
<<<<<<< HEAD
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	//This func is for server API
	r := gin.Default()
	//Method for get post and login
	r.GET("/cekuser", controllers.CekUser)
	r.GET("/post", controllers.GetPost)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
=======
	// if you local change port using your available port
	port := os.Getenv("PORT")
	//This func is for server API
	r := gin.Default()
	//Method for get post and login
	r.GET("/post", controllers.GetPost)
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
	r.Use(middleware.CheckToken)
	//Method for post
	r.POST("/post", controllers.CreatePost)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)
	//Method for logout
	r.GET("/logout", controllers.Logout)
	//Method for profile
<<<<<<< HEAD
	r.GET("/profile/", controllers.GetProfile)
	r.PUT("profile/", controllers.UpdateProfile)
=======
	r.GET("/profile/:id", controllers.GetProfile)
	r.PUT("profile/:id", controllers.UpdateProfile)
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
	r.Run(":" + port)
}
