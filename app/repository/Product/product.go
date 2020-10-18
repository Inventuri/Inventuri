package Product

import (
	"github.com/jinzhu/gorm"
)

func NewSqlProductRepo(db *gorm.DB) *sqlProductRepo {
	return &sqlProductRepo{Database: db}
}

type sqlProductRepo struct {
	Database *gorm.DB
}
