package Business

import "Inventuri/app/models"

func (db sqlBusinessRepo) GetByID(id uint) (*models.Business, error) {
	business := models.Business{}
	err := db.Database.Model(&business).Where("id = ?", id).Find(&business).Error
	return &business, err
}
