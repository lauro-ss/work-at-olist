package controllers_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lauro-ss/work-at-olist/internal/controllers"
	"github.com/lauro-ss/work-at-olist/internal/data"
)

func Setup() (*gin.Engine, error) {
	r := gin.Default()
	//TODO: mock here
	r.GET("/author", controllers.ListAuthors(nil))

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
