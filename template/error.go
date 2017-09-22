package template

<<<<<<< HEAD
import "github.com/rezandry/kejarmimpi/models"
=======
import "kejarmimpi/models"
>>>>>>> 9fafbf152d3328feb0d1dc7939112e8a940716f4

//Response is func for template Error message
func Response(res *models.Response) map[string]interface{} {

	data := map[string]interface{}{
		"code":    res.Code,
		"message": res.Message,
		"token":   res.Token,
	}
	return data
}
