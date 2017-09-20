package template

import "kejarmimpi/models"

//Response is func for template Error message
func Response(res *models.Response) map[string]interface{} {

	data := map[string]interface{}{
		"code":    res.Code,
		"message": res.Message,
		"token":   res.Token,
	}
	return data
}