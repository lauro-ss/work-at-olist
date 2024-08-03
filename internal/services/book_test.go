package services_test

import (
	"testing"

	"github.com/lauro-ss/work-at-olist/internal/data"
	"github.com/lauro-ss/work-at-olist/internal/services"
)

func TestList(t *testing.T) {
	db, err := data.OpenAndMigrate("user=postgres password=postgres host=localhost port=5432 database=postgres")
	if err != nil {
		t.Fatal(err)
	}
	br := services.NewBookRepository(db)
	_, err = br.List()
	if err != nil {
		t.Fatal(err)
	}
}
