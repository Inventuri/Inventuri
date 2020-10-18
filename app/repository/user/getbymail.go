package user

import "Inventuri/app/models"

func (db sqlUserRepo) GetByMail(mail string) (*models.User, error) {
	user := models.User{}
	err := db.Database.Model(&user).Where("email = ?", mail).Find(&user).Error
	return &user, err
}
