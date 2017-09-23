package template

import "github.com/rezandry/kejarmimpi/models"

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
