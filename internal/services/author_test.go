package services_test

import (
	"testing"

	"github.com/lauro-ss/work-at-olist/internal/data"
	"github.com/lauro-ss/work-at-olist/internal/services"
)

func TestCreateBatch(t *testing.T) {
	db, err := data.OpenAndMigrate("user=postgres password=postgres host=localhost port=9432 database=postgres")
	if err != nil {
		t.Fatal(err)
	}
	ar := services.NewAuthorRepository(db)

	_, err = ar.CreateBatch([]data.Author{{Name: "T1"}, {Name: "T2"}})
	if err != nil {
		t.Fatal(err)
	}
}
