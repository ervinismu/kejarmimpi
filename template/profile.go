package template

<<<<<<< HEAD
import "github.com/rezandry/kejarmimpi/models"
=======
import "kejarmimpi/models"
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4

//Profile is making template for get data profile
func Profile(user *models.User) map[string]interface{} {

	data := map[string]interface{}{
		"id":        user.ID,
		"name":      user.Name,
		"photo":     user.Photo,
		"job":       user.Job,
		"address":   user.Address,
		"status":    user.Status,
		"biograph":  user.Biograph,
		"education": user.Education,
		"career":    user.Career,
	}
	return data
}
