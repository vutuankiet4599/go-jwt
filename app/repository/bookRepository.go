package repository

import (
	"github.com/vutuankiet4599/go-jwt/app/models"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

type BookRepository interface {
	GetAll() (*gorm.DB ,*[]models.Book)
	GetOne(id uint) (*gorm.DB, *models.Book)
	Insert(title string, page int, userId uint) (*gorm.DB, *models.Book)
	Update(id uint, title string, page int) (*gorm.DB, *models.Book)
	DeleteOneById(id uint) (*gorm.DB)
	DeleteAll(userId uint) (*gorm.DB)
}

func (r *bookRepository) GetAll() (*gorm.DB, *[]models.Book) {
	var books []models.Book
	response := r.db.Find(&books)
	return response, &books
}

func (r *bookRepository) GetOne(id uint) (*gorm.DB, *models.Book) {
	var book models.Book
	response := r.db.First(&book, id)
	return response, &book
}

func (r *bookRepository) Insert(title string, page int, userId uint) (*gorm.DB, *models.Book) {
	book := models.Book{
		Title: title,
		Page: page,
		UserId: userId,
	}
	response := r.db.Create(&book)
	return response, &book
}

func (r *bookRepository) Update(id uint, title string, page int) (*gorm.DB, *models.Book) {
	var book models.Book
	response := r.db.First(&book, id)
	if response.Error != nil {
		return response, &models.Book{}
	}
	response = r.db.Model(&book).Updates(models.Book{
		Title: title,
		Page: page,
	})
	return response, &book
}

func (r *bookRepository) DeleteOneById(id uint) (*gorm.DB) {
	response := r.db.Delete(&models.Book{}, id)
	return response
}

func (r *bookRepository) DeleteAll(userId uint) (*gorm.DB) {
	response := r.db.Delete(&models.Book{}, &models.Book{
		UserId: userId,
	})
	return response
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}
