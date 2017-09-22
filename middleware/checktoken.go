package middleware

import (
<<<<<<< HEAD
	"github.com/gin-gonic/gin"
	"github.com/rezandry/kejarmimpi/controllers"
	"github.com/rezandry/kejarmimpi/models"
=======
	"kejarmimpi/configs"
	"kejarmimpi/models"

	"github.com/gin-gonic/gin"
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
)

// CheckToken is func for checking token
// Logic checktoken :
// 1. check if token empty
// 2. check if token is not same in database
func CheckToken(c *gin.Context) {
<<<<<<< HEAD
	db := controllers.InitDb()
=======
	db := configs.InitDb()
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
	defer db.Close()

	var user models.User

	token := c.Request.Header.Get("token")
<<<<<<< HEAD
=======

>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4
	//1. check if token empty
	if token == "" {
		data := map[string]interface{}{
			"code":    "402",
			"message": "token API required",
		}
		c.JSON(402, data)
		c.Abort()
		return
	}

	//2. check if token not same as database
	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
		data := map[string]interface{}{
			"code":    "401",
			"message": "token API invalid",
		}
		c.JSON(401, data)
		c.Abort()
		return
	}

	c.Next()

}
