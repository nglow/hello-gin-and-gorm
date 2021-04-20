package service

import (
	"fmt"
	"github.com/mashingan/smapping"
	"helloGinAndGorm/dto"
	"helloGinAndGorm/entity"
	"helloGinAndGorm/repository"
	"log"
)

// BookService is a ......
type BookService interface {
	Insert(b dto.BookCreateDto) entity.Book
	Update(b dto.BookUpdateDto) entity.Book
	Delete(b entity.Book)
	All() []entity.Book
	FindById(bookId uint64) entity.Book
	IsAllowedToEdit(userId string, bookId uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService {
		bookRepository: bookRepo,
	}
}

func (service *bookService) Insert(b dto.BookCreateDto) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.bookRepository.InsertBook(book)
	return res
}

func (service *bookService) Update(b dto.BookUpdateDto) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.bookRepository.UpdateBook(book)
	return res
}

func (service *bookService) Delete(b entity.Book) {
	service.bookRepository.DeleteBook(b)
}

func (service *bookService) All() []entity.Book {
	return service.bookRepository.AllBooks()
}

func (service *bookService) FindById(bookId uint64) entity.Book {
	return service.bookRepository.FindBookById(bookId)
}

func (service *bookService) IsAllowedToEdit(userId string, bookId uint64) bool {
	b := service.bookRepository.FindBookById(bookId)
	id := fmt.Sprintf("%v", b.UserId)
	return userId == id
}