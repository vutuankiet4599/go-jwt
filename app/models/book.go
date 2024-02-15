package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string `json:"title"`
	Page int `json:"page"`
	UserId uint `json:"userId"`
	User User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
