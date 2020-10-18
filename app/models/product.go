package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Product struct {
	ID uint `gorm:"primaryKey"`

	Name string
	Amount int
	MinAmount int
	MaxAmount int
	NetPrice decimal.Decimal `gorm:"type:numeric"`
	taxGroup int
	SellingPrice decimal.Decimal `gorm:"type:numeric"`
	EANCode string
	Storagegroup string
	Storageplace string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}