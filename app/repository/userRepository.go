package repository

import (
	"github.com/vutuankiet4599/go-jwt/app/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	IsDuplicateEmail(email string) (*gorm.DB)
	CreateUser(email, password, name string) (*gorm.DB, *models.User)
	GetUser(id uint) (*gorm.DB, *models.User)
	VerifyCredentials(email string) (*gorm.DB, *models.User)
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user models.User
	return r.db.Where("email = ?", email).Take(&user)
}

func (r *userRepository) CreateUser(email, password, name string) (*gorm.DB, *models.User) {
	var user models.User = models.User{
		Email: email,
		Password: password,
		Name: name,
	}
	response := r.db.Create(&user)
	return response, &user
}

func (r *userRepository) GetUser(id uint) (tx *gorm.DB, user *models.User) {
	var data models.User
	response := r.db.Where("id = ?", id).Take(&data)
	return response, &data
}

func (r *userRepository) VerifyCredentials(email string) (*gorm.DB, *models.User) {
	var user models.User
	response := r.db.Preload("Books").Where(&models.User{
		Email: email,
	}).Take(&user)
	return response, &user
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
