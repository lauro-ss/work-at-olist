package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lauro-ss/work-at-olist/docs"
	"github.com/lauro-ss/work-at-olist/internal/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	r.GET("/author", controllers.ListAuthors())

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
