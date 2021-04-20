package repository

import (
	"gorm.io/gorm"
	"helloGinAndGorm/entity"
)

// BookRepository is a ....
type BookRepository interface {
	InsertBook(b entity.Book) entity.Book
	UpdateBook(b entity.Book) entity.Book
	DeleteBook(b entity.Book)
	AllBooks() []entity.Book
	FindBookById(bookId uint64) entity.Book
}

type bookConnection struct {
	connection *gorm.DB
}

// NewBookRepository creates a instance BookRepository
func NewBookRepository(dbConn *gorm.DB) BookRepository {
	return &bookConnection {
		connection:dbConn,
	}
}

func (db *bookConnection) InsertBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bookConnection) UpdateBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bookConnection) DeleteBook(b entity.Book) {
	db.connection.Delete(&b)
}

func (db *bookConnection) FindBookById(bookId uint64) entity.Book {
	var book entity.Book
	db.connection.Preload("User").Find(&book, bookId)
	return book
}

func (db *bookConnection) AllBooks() []entity.Book {
	var books []entity.Book
	db.connection.Preload("User").Find(&books)
	return books
}