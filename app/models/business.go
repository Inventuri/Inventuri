package models

import "time"

type Business struct {
	ID uint `gorm:"primaryKey"`

	Name string
	Description string
	Street string
	City string
	Zip string
	Phone string

	UserID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
