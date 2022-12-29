package models

import (
	"errors"

	"github.com/google/uuid"
)

type Book struct {
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

type Table struct {
	Name string `json:"name"`
}

var _books = []Book{
	{Uuid: uuid.New().String(), Name: "raxit", Publication: "X"},
	{Uuid: uuid.New().String(), Name: "Prisha", Publication: "X"},
}

func GetAll() ([]Book, error) {
	return _books, nil
}

func Get(uuid string) (Book, error) {
	for _, book := range _books {
		if book.Uuid == uuid {
			return book, nil
		}
	}

	return Book{}, errors.New("Book not found")
}

func Create(book *Book) error {
	book.Uuid = uuid.New().String()
	_books = append(_books, *book)
	return nil
}

func Update(book Book) (Book, error) {
	isFound := false

	for i, obj := range _books {
		if obj.Uuid == book.Uuid {
			// obj = book
			_books = append(_books[:i], book)

			isFound = true
		}
	}

	if !isFound {
		return Book{}, errors.New("Book not found")
	}

	return book, nil
}
