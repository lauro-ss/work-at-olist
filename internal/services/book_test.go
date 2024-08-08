package services_test

import (
	"testing"

	"github.com/lauro-ss/work-at-olist/internal/data"
	"github.com/lauro-ss/work-at-olist/internal/services"
)

func Setup() (*services.BookRepository, error) {
	db, err := data.OpenAndMigrate("user=postgres password=postgres host=localhost port=5432 database=postgres")
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
	book := br.Create(data.Book{Name: "Clean Code", Edition: 1, PublicationYear: 2012})
	if book.Id == 0 {
		t.Error("Expected Id Value, got", book.Id)
	}
}
