package Business

import "Inventuri/app/models"

func (db sqlBusinessRepo) Create(p *models.Business) (*models.Business, error) {
	business := models.Business{}
	err := db.Database.Create(p).Scan(&business).Error
	return &business, err
}
