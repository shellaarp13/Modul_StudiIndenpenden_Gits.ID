package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm: "size:36; not null; uniqueIndex;primary_key`
	Addresses     []Address
	FisrtName     string `gorm: "size:100;not null`
	LastName      string `gorm: "size:100;not null`
	Email         string `gorm: "size:100;not null`
	Password      string `gorm: "size:100;not null; uniqueIndex`
	RememberToken string `gorm: "size:100;not null`
	CreateAt      time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
