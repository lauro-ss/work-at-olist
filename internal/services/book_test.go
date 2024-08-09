package services_test

import (
	"testing"

	"github.com/lauro-ss/work-at-olist/internal/data"
	"github.com/lauro-ss/work-at-olist/internal/services"
)

func Setup() (*services.BookRepository, error) {
	db, err := data.OpenAndMigrate("user=postgres password=postgres host=localhost port=9432 database=postgres")
	if err != nil {
		return nil, err
	}
	return services.NewBookRepository(db), nil
}

func TestList(t *testing.T) {
	br, err := Setup()
	if err != nil {
		t.Fatal(err)
	}
	_, err = br.List()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreate(t *testing.T) {
	br, err := Setup()
	if err != nil {
		t.Fatal(err)
	}
	book, err := br.Create(data.Book{Name: "Clean Code", Edition: 1, PublicationYear: 2012})
	if err != nil {
		t.Fatal(err)
	}
	if book.Id == 0 {
		t.Error("Expected Id Value, got", book.Id)
	}
}

func TestCreateBookAuthor(t *testing.T) {
	br, err := Setup()
	if err != nil {
		t.Fatal(err)
	}
	ar := services.NewAuthorRepository(br.Db)
	authors, err := ar.CreateBatch([]data.Author{{Name: "Lauro"}})
	if err != nil {
		t.Fatal(err)
	}
	book, err := br.Create(data.Book{Name: "LOL", Edition: 1, PublicationYear: 2012, Authors: authors})
	if err != nil {
		t.Fatal(err)
	}
	if book.Id == 0 {
		t.Error("Expected Id Value, got", book.Id)
	}
}

func TestGet(t *testing.T) {
	br, err := Setup()
	if err != nil {
		t.Fatal(err)
	}
	book, err := br.Create(data.Book{Name: "Golang for Noobs", Edition: 1, PublicationYear: 2012})
	if err != nil {
		t.Fatal(err)
	}
	if book.Id == 0 {
		t.Error("Expected Id Value, got", book.Id)
	}

	if br.Get(book.Id) == nil {
		t.Error("Expected Get a Book, got nil")
	}
}
