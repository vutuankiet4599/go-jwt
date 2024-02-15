package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `json:"email" gorm:"uniqueIndex"`
	Name string `json:"name"`
	Password string `json:"-"`
	Books []Book `json:"books"`
}
