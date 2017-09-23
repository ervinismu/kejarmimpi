package controllers

import "github.com/rezandry/kejarmimpi/models"

//CountCollabs is for count how many people collabs
func CountCollabs(idPost uint) (count int) {
	db := InitDb()
	defer db.Close()

	var collabsCount int
	var collabs []models.Collabs

	if err := db.Select("id").Where("id_post = ?", idPost).Find(&collabs).Error; err != nil {
		collabsCount = 0
	} else {
		collabsCount = len(collabs)
	}
	return collabsCount
}
