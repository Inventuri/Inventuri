package Product

import "Inventuri/app/models"

func (db sqlProductRepo) Delete(id uint) error {
	product := models.Product{ID: id}
	err := db.Database.Delete(&product).Error
	return err
}