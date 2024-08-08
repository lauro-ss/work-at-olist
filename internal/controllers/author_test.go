package controllers_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lauro-ss/work-at-olist/internal/controllers"
	"github.com/lauro-ss/work-at-olist/internal/data"
	"github.com/lauro-ss/work-at-olist/internal/services"
)

func Setup() (*gin.Engine, error) {
	db, err := data.OpenAndMigrate("user=postgres password=postgres host=localhost port=5432 database=postgres")
	if err != nil {
		return nil, err
	}
	ar := services.NewAuthorRepository(db)

	r := gin.Default()
	r.GET("/author", controllers.ListAuthors(ar))

	return r, nil
}

func TestGetAuthor(t *testing.T) {
	r, err := Setup()
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, httptest.NewRequest("GET", "/author", nil))
	if recorder.Code != 200 {
		t.Error("Expected 200, got", recorder.Code)
	}
	var authors []data.Author
	err = json.NewDecoder(recorder.Body).Decode(&authors)
	if err != nil {
		t.Fatal(err)
	}
}
