package user

import "github.com/jinzhu/gorm"

func NewSqlUserRepo(db *gorm.DB) *sqlUserRepo {
	return &sqlUserRepo{Database: db}
}

type sqlUserRepo struct {
	Database *gorm.DB
}
