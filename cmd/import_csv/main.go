package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/lauro-ss/work-at-olist/internal/data"
	"github.com/lauro-ss/work-at-olist/internal/services"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		return
	}
	file, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	//skip column name
	_, err = csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}

	row, err := csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}

	var authors []data.Author
	for row != nil {
		authors = append(authors, data.Author{Name: row[0]})
		row, err = csvReader.Read()
		if err != nil && !errors.Is(err, io.EOF) {
			log.Fatal(err)
		}
	}

	db, err := data.OpenAndMigrate("user=postgres password=postgres host=localhost port=5432 database=postgres")
	if err != nil {
		log.Fatal(err)
	}

	authors, err = services.NewAuthorRepository(db).CreateBatch(authors)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted %v authors, last id: %v\n", len(authors), authors[len(authors)-1].Id)
}
