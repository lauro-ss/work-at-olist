package services

import "github.com/lauro-ss/work-at-olist/internal/data"

type AuthorRepository struct {
	Db *data.Database
}

func NewAuthorRepository(Db *data.Database) *AuthorRepository {
	return &AuthorRepository{
		Db: Db,
	}
}

func (ar *AuthorRepository) List() (authors []data.Author) {
	ar.Db.Select(ar.Db.Author).Scan(&authors)
	return authors
}

func (ar *AuthorRepository) CreateBatch(author []data.Author) ([]data.Author, error) {
	_, err := ar.Db.Insert(ar.Db.Author).Value(&author)
	return author, err
}
