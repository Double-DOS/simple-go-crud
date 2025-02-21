package models

import (
	"github.com/Double-DOS/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	// db.NewRecord(b)
	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.First(&getBook, "ID=?", Id)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Delete(&book, Id)
	return book
}

func UpdateBook(Id int64, data *Book) Book {
	var book Book
	db.Model(book).Where("ID = ?", Id).Updates(data)
	db.First(&book, "ID=?", Id)
	return book
}
