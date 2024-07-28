package data

import (
	"github.com/lauro-ss/goe"
	"github.com/lauro-ss/postgres"
)

type Author struct {
	Id    uint
	Name  string `goe:"type:varchar(50)"`
	Books []Book `goe:"table:BookAuthor"`
}

type Book struct {
	Id              uint
	Name            string `goe:"type:varchar(50)"`
	Edition         uint8
	PublicationYear uint16
	Authors         []Author `goe:"table:BookAuthor"`
}

type Database struct {
	Author *Author
	Book   *Book
	*goe.DB
}

func OpenAndMigrate(dns string) (*Database, error) {
	db := &Database{DB: &goe.DB{}}
	err := goe.Open(db, postgres.Open(dns))
	if err != nil {
		return nil, err
	}

	err = db.Migrate(goe.MigrateFrom(db))
	if err != nil {
		return nil, err
	}
	return db, nil
}
