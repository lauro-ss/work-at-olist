package services

import (
	"fmt"

	"github.com/lauro-ss/work-at-olist/internal/data"
)

type BookRepository struct {
	Db *data.Database
}

func NewBookRepository(Db *data.Database) *BookRepository {
	return &BookRepository{
		Db: Db,
	}
}

func (br *BookRepository) List() (books []data.Book, err error) {
	_, err = br.Db.Select(br.Db.Book).Scan(&books)
	return books, err
}

func (br *BookRepository) Get(id uint) *data.Book {
	Db := br.Db

	var book data.Book
	Db.Select(Db.Book).
		Where(Db.Equals(&Db.Book.Id, id)).
		Scan(&book)

	if book.Id != 0 {
		return &book
	}
	return nil
}

func (br *BookRepository) Create(book data.Book) (*data.Book, error) {
	_, err := br.Db.Insert(br.Db.Book).Value(&book)
	if err != nil {
		return nil, err
	}
	if len(book.Authors) > 0 {
		bookAuthor := make([]uint, len(book.Authors)*2)
		c := 0
		for i := range book.Authors {
			bookAuthor[i+c], bookAuthor[i+c+1] = book.Id, book.Authors[i].Id
			c++
		}
		_, err = br.Db.InsertIn(br.Db.Book, br.Db.Author).Values(bookAuthor)
		if err != nil {
			return nil, err
		}
	}
	return &book, nil
}

func (br *BookRepository) Update(id uint, book data.Book) data.Book {
	Db := br.Db
	Db.Update(Db.Book).
		Where(Db.Equals(&Db.Book.Id, id)).
		Value(&book)

	if len(book.Authors) > 0 {
		bookAuthor := make([]uint, len(book.Authors)*2)
		c := 0
		for i := range book.Authors {
			bookAuthor[i+c], bookAuthor[i+c+1] = book.Id, book.Authors[i].Id
			c++
		}
		Db.DeleteIn(Db.Book, Db.Author).Where(Db.Equals(&Db.Book.Id, id))
		Db.InsertIn(Db.Book, Db.Author).Values(bookAuthor)
	}
	return book
}

func (br *BookRepository) Delete(id uint) bool {
	Db := br.Db

	_, err := Db.DeleteIn(Db.Book, Db.Author).Where(Db.Equals(&Db.Book.Id, id))
	if err != nil {
		fmt.Println(err)
		return false
	}

	_, err = Db.Delete(Db.Book).
		Where(Db.Equals(&Db.Book.Id, id))
	return err == nil
}
