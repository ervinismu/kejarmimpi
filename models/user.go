package models

//User is struct for database User
type User struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	Photo           string `json:"photo"`
	Job             string `json:"job"`
	Address         string `json:"address"`
	Status          string `json:"status"`
	Biograph        string `json:"biograph"`
	Education       string `json:"education"`
	Career          string `json:"career"`
	Token           string `json:"token"`
}
