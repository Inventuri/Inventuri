package Business

import "Inventuri/app/models"

func (db sqlBusinessRepo) Delete(id uint) error {
	business := models.Business{ID: id}
	err := db.Database.Delete(&business).Error
	return err
}
