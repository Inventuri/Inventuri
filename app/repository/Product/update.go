package Product

import "Inventuri/app/models"

func (db sqlProductRepo) Update(p *models.Product) (*models.Product, error) {
	product := models.Product{}
	err := db.Database.Model(&product).Update(p).Scan(&product).Error
	return &product, err
}
