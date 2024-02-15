package service

import (
	"github.com/vutuankiet4599/go-jwt/app/models"
	"github.com/vutuankiet4599/go-jwt/app/repository"
	"github.com/vutuankiet4599/go-jwt/app/request"
)

type bookService struct {
	bookRepository repository.BookRepository
}

type BookService interface {
	GetAll() (*[]models.Book, bool, string)
	GetOneById(id uint) (*models.Book, bool, string)
	Insert(bookData *request.InsertBookRequest, userId uint) (*models.Book, bool, string)
	Update(bookData *request.UpdateBookRequest, id uint) (*models.Book, bool, string)
	DeleteOneById(id uint) (bool, string)
	DeleteAll(userId uint) (bool, string)
}

func (s *bookService) GetAll() (*[]models.Book, bool, string) {
	response, books := s.bookRepository.GetAll()
	if response.Error != nil {
		return &[]models.Book{}, true, response.Error.Error() 
	}
	return books, false, ""
}

func (s *bookService) GetOneById(id uint) (*models.Book, bool, string) {
	response, book := s.bookRepository.GetOne(id)
	if response.Error != nil {
		return &models.Book{}, true, response.Error.Error()
	}
	return book, false, ""
}

func (s *bookService) Insert(bookData *request.InsertBookRequest, userId uint) (*models.Book, bool, string) {
	response, insertedBook := s.bookRepository.Insert(bookData.Title, bookData.Page, userId)
	if response.Error != nil {
		return &models.Book{}, true, response.Error.Error()
	}
	return insertedBook, false, ""
}

func (s *bookService) Update(bookData *request.UpdateBookRequest, id uint) (*models.Book, bool, string) {
	response, updatedBook := s.bookRepository.Update(id, bookData.Title, bookData.Page)
	if response.Error != nil {
		return &models.Book{}, true, response.Error.Error()
	}
	return updatedBook, false, ""
}

func (s *bookService) DeleteOneById(id uint) (bool, string) {
	response := s.bookRepository.DeleteOneById(id)
	if response.Error != nil {
		return true, response.Error.Error()
	}
	return false, ""
}

func (s *bookService) DeleteAll(userId uint) (bool, string) {
	response := s.bookRepository.DeleteAll(userId)
	if response.Error != nil {
		return true, response.Error.Error()
	}
	return false, ""
}
func NewBookService(bookRepository repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepository,
	}
}
