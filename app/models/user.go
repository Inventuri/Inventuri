package models

import "time"

type User struct {
	ID uint `gorm:"primaryKey"`

	Name string
	Surname string
	Mail string
	PasswordHashed string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
