package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lauro-ss/work-at-olist/docs"
	"github.com/lauro-ss/work-at-olist/internal/controllers"
	"github.com/lauro-ss/work-at-olist/internal/data"
	"github.com/lauro-ss/work-at-olist/internal/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	db, err := data.OpenAndMigrate("user=postgres password=postgres host=localhost port=5432 database=postgres")
	if err != nil {
		log.Fatal(err)
	}
	ar := services.NewAuthorRepository(db)
	br := services.NewBookRepository(db)

	r.GET("/author", controllers.ListAuthors(ar))

	r.GET("/book", controllers.ListBooks(br))
	r.GET("/book/:id", controllers.GetBook(br))
	r.POST("/book", controllers.CreateBook(br))
	r.PUT("/book", controllers.UpdateBook(br))
	r.DELETE("/book/:id", controllers.DeleteBook(br))

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
