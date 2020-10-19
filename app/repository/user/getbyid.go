package user

import "Inventuri/app/models"

func (db sqlUserRepo) GetByID(id uint) (*models.User, error) {
	user := models.User{}
	err := db.Database.Model(&user).Where("id = ?", id).Find(&user).Error
	return &user, err
}
