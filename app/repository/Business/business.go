package Business

import (
	"github.com/jinzhu/gorm"
)

func NewSQLBusinessRepo(db *gorm.DB) *sqlBusinessRepo {
	return &sqlBusinessRepo{Database: db}
}

type sqlBusinessRepo struct {
	Database *gorm.DB
}
