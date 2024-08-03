package services

import (
	"fmt"

	"github.com/lauro-ss/work-at-olist/internal/data"
)

type BookRepository struct {
	db *data.Database
}

func NewBookRepository(db *data.Database) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (br *BookRepository) List() (books []data.Book, err error) {
	_, err = br.db.Select(br.db.Book).Scan(&books)
	return books, err
}

func (br *BookRepository) Get(id uint) *data.Book {
	db := br.db

	var book data.Book
	db.Select(db.Book).
		Where(db.Equals(&db.Book.Id, id)).
		Scan(&book)

	if book.Id != 0 {
		return &book
	}
	return nil
}

func (br *BookRepository) Create(book data.Book) data.Book {
	br.db.Insert(br.db.Book).Value(&book)
	fmt.Println(book)
	if len(book.Authors) > 0 {
		bookAuthor := make([]uint, len(book.Authors)*2)
		c := 0
		for i := range book.Authors {
			bookAuthor[i+c], bookAuthor[i+c+1] = book.Id, book.Authors[i].Id
			c++
		}
		br.db.InsertIn(br.db.Book, br.db.Author).Values(bookAuthor)
	}
	return book
}

func (br *BookRepository) Update(id uint, book data.Book) data.Book {
	db := br.db
	db.Update(db.Book).
		Where(db.Equals(&db.Book.Id, id)).
		Value(&book)

	if len(book.Authors) > 0 {
		bookAuthor := make([]uint, len(book.Authors)*2)
		c := 0
		for i := range book.Authors {
			bookAuthor[i+c], bookAuthor[i+c+1] = book.Id, book.Authors[i].Id
			c++
		}
		db.DeleteIn(db.Book, db.Author).Where(db.Equals(&db.Book.Id, id))
		db.InsertIn(db.Book, db.Author).Values(bookAuthor)
	}
	return book
}

func (br *BookRepository) Delete(id uint) bool {
	db := br.db

	_, err := db.DeleteIn(db.Book, db.Author).Where(db.Equals(&db.Book.Id, id))
	if err != nil {
		fmt.Println(err)
		return false
	}

	_, err = db.Delete(db.Book).
		Where(db.Equals(&db.Book.Id, id))
	return err == nil
}
