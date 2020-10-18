package user

import "Inventuri/app/models"

func (db sqlUserRepo) GetById(id uint) (*models.User, error) {
	user := models.User{}
	err := db.Database.Model(&user).Where("id = ?", id).Find(&user).Error
	return &user, err
}
