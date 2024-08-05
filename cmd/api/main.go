package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lauro-ss/work-at-olist/docs"
	"github.com/lauro-ss/work-at-olist/internal/controllers"
	"github.com/lauro-ss/work-at-olist/internal/data"
	"github.com/lauro-ss/work-at-olist/internal/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	dns := os.Getenv("DATABASE_DNS")
	if dns == "" {
		dns = "user=postgres password=postgres host=localhost port=5432 database=postgres"
	}
	db, err := data.OpenAndMigrate(dns)
	if err != nil {
		log.Fatal(err)
	}
	ar := services.NewAuthorRepository(db)
	br := services.NewBookRepository(db)

	r := gin.Default()
	r.GET("/author", controllers.ListAuthors(ar))

	r.GET("/book", controllers.ListBooks(br))
	r.GET("/book/:id", controllers.GetBook(br))
	r.POST("/book", controllers.CreateBook(br))
	r.PUT("/book", controllers.UpdateBook(br))
	r.DELETE("/book/:id", controllers.DeleteBook(br))

	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(os.Getenv("SERVER_HOST") + os.Getenv("SERVER_PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
