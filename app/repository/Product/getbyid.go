package Product

import "Inventuri/app/models"

func (db sqlProductRepo) GetById(id uint) (*models.Product, error) {
	product := models.Product{}
	err := db.Database.Model(&product).Where("id = ?", id).Find(&product).Error
	return &product, err
}