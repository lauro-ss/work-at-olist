package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lauro-ss/work-at-olist/internal/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controllers.ListAuthors())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
