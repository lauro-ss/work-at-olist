package services

import "github.com/lauro-ss/work-at-olist/internal/data"

type AuthorRepository struct {
	db *data.Database
}

func NewAuthorRepository(db *data.Database) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (ar *AuthorRepository) List() (authors []data.Author) {
	ar.db.Select(ar.db.Author).Scan(&authors)
	return authors
}

func (ar *AuthorRepository) CreateBatch(author []data.Author) ([]data.Author, error) {
	_, err := ar.db.Insert(ar.db.Author).Value(&author)
	return author, err
}
