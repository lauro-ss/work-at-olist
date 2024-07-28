package services

import "github.com/lauro-ss/work-at-olist/internal/data"

type AuthorRepository struct {
	*data.Database
}

func newAuthorRepository(db *data.Database) *AuthorRepository {
	return &AuthorRepository{
		Database: db,
	}
}

func (ar *AuthorRepository) List() (authors []data.Author) {
	ar.Select(ar.Author).Scan(&authors)
	return authors
}

func (ar *AuthorRepository) CreateBatch(author []data.Author) []data.Author {
	ar.Insert(ar.Author).Value(&author)
	return author
}
