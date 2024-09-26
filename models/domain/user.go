package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	UserName string
	Password string
}
