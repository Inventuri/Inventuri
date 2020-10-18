package user

import "Inventuri/app/models"

func (db sqlUserRepo) Update(u *models.User) (*models.User, error) {
	user := models.User{}
	err := db.Database.Model(&user).Update(u).Scan(&user).Error
	return &user, err
}
