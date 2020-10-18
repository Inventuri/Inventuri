package Product

import "Inventuri/app/models"

func (db sqlProductRepo) Create(p *models.Product) (*models.Product, error) {
	product := models.Product{}
	err := db.Database.Create(p).Scan(&product).Error
	return &product, err
}
