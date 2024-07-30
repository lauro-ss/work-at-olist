package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lauro-ss/work-at-olist/internal/services"
)

// @Summary			Book
//
// @Accept			json
// @Produce			json
// @Success			200	{object}	[]data.Book
// @Router			/book [get]
func ListBooks(br *services.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, br.List())
	}
}

// @Summary			Book
//
// @Accept			json
// @Produce			json
// @Param			id	path		uint	true	"Book ID"
// @Success			200	{object}	data.Book
// @Failure			404
// @Router			/book/{id} [get]
func GetBook(br *services.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message":     "not found",
				"status_code": "404",
			})
		}
		b := br.Get(uint(id))
		if b != nil {
			c.JSON(http.StatusOK, b)
			return
		}
		c.JSON(http.StatusNotFound, gin.H{
			"message":     "not found",
			"status_code": "404",
		})
	}
}
