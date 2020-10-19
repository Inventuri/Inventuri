package Product

import "Inventuri/app/models"

func (db sqlProductRepo) GetByID(id uint) (*models.Product, error) {
	product := models.Product{}
	err := db.Database.Model(&product).Where("id = ?", id).Find(&product).Error
	return &product, err
}