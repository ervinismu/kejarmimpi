package models

// Collabs is struct for save collabs
type Collabs struct {
	ID      uint   `form:"id" binding:"required"`
	IDPost uint   `form:"idPost" binding:"required"`
	IDUPost uint   `form:"idUPost" binding:"required"`
	IDUser  uint   `form:"idUser" binding:"required"`
	Check   bool   `form:"check" binding:"required"`
	Message string `form:"message" binding:"required"`
}
