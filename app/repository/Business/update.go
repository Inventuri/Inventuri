package Business

import "Inventuri/app/models"

func (db sqlBusinessRepo) Update(p *models.Business) (*models.Business, error) {
	business := models.Business{}
	err := db.Database.Model(&business).Update(p).Scan(&business).Error
	return &business, err
}
