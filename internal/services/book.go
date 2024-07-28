package services

import "github.com/lauro-ss/work-at-olist/internal/data"

type BookRepository struct {
	*data.Database
}

func newBookRepository(db *data.Database) *BookRepository {
	return &BookRepository{
		Database: db,
	}
}

func (br *BookRepository) List() (books []data.Book) {
	br.Select(br.Author).Scan(&books)
	return books
}

func (br *BookRepository) CreateBatch(book data.Book) data.Book {
	br.Insert(br.Author).Value(&book)
	return book
}
