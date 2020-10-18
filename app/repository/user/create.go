package user

import "Inventuri/app/models"

func (db sqlUserRepo) Create(u *models.User) (*models.User, error) {
	user := models.User{}
	err := db.Database.Create(u).Scan(&user).Error
	return &user, err
}
