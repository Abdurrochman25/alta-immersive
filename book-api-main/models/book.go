package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string
	Author string
}

type GormBookModel struct {
	db *gorm.DB
}

func NewBookModel(db *gorm.DB) *GormBookModel {
	return &GormBookModel{db: db}
}

type BookModel interface {
	GetAll() ([]Book, error)
	Get(bookId int) (Book, error)
	Insert(Book) (Book, error)
	Edit(book Book, bookId int) (Book, error)
	Delete(bookId int) (Book, error)
}

func (m *GormBookModel) GetAll() ([]Book, error) {
	var book []Book
	if err := m.db.Find(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (m *GormBookModel) Get(bookId int) (Book, error) {
	var book Book
	if err := m.db.Find(&book, bookId).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (m *GormBookModel) Insert(book Book) (Book, error) {
	if err := m.db.Save(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (m *GormBookModel) Edit(newBook Book, bookId int) (Book, error) {
	var book Book
	if err := m.db.Find(&book, bookId).Error; err != nil {
		return book, err
	}

	book.Title = newBook.Title
	book.Author = newBook.Author

	if err := m.db.Save(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (m *GormBookModel) Delete(bookId int) (Book, error) {
	var book Book
	if err := m.db.Find(&book, bookId).Error; err != nil {
		return book, err
	}

	if err := m.db.Delete(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}
