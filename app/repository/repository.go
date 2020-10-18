package repository

import "Inventuri/app/models"

type ProductRepo interface {
	Create(p *models.Product) (*models.Product, error)
	Update(p *models.Product) (*models.Product, error)
	Delete(id uint) error
	GetById(id uint) (*models.Product, error)
}

type UserRepo interface {
	Create(u *models.User) (*models.User, error)
	Update(u *models.User) (*models.User, error)
	Delete(id uint) error
	GetById(id uint) (*models.User, error)
	GetByMail(mail string) (*models.User, error)
}

type BusinessRepo interface {
	Create(u *models.Business) (*models.Business, error)
	Update(u *models.Business) (*models.Business, error)
	Delete(id uint) error
	GetById(id uint) (*models.User, error)
}