package models

//Star is used for artikel
type Star struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Job     string `json:"job"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
