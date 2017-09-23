package models

//User is struct for database User
type User struct {
	ID              uint   `form:"id" json:"id"`
	Name            string `form:"name" json:"name"`
	Email           string `form:"email" json:"email"`
	Password        string `form:"password" json:"password"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm"`
	Photo           string `form:"photo" json:"photo"`
	Job             string `form:"job" json:"job"`
	Address         string `form:"address" json:"address"`
	Status          string `form:"status" json:"status"`
	Biograph        string `form:"biograph" json:"biograph"`
	Education       string `form:"education" json:"education"`
	Career          string `form:"career" json:"career"`
	Token           string `form:"token" json:"token"`
}
