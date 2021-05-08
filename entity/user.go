package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Identifier string
	Secret string
}
