package user

import "Inventuri/app/models"

func (db sqlUserRepo) Delete(id uint) error {
	user := models.User{ID: id}
	err := db.Database.Delete(&user).Error
	return err
}
