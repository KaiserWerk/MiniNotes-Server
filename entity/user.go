package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Secret string
	Content string
}
