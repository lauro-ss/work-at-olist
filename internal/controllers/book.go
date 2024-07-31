package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lauro-ss/work-at-olist/internal/data"
	"github.com/lauro-ss/work-at-olist/internal/services"
)

// @Summary			List Book
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

// @Summary			Get Book
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

// @Summary			Create Book
//
// @Accept			json
// @Produce			json
// @Param			book	body		data.Book	true "Book JSON"
// @Success			200	{object}	data.Book
// @Failure			400
// @Router			/book [post]
func CreateBook(br *services.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var b data.Book

		err := c.BindJSON(&b)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":     "bad request",
				"status_code": "400",
			})
		}
		c.JSON(http.StatusOK, br.Create(b))
	}
}

// @Summary			Update Book
//
// @Accept			json
// @Produce			json
// @Param			book	body		data.Book	true "Book JSON"
// @Success			200	{object}	data.Book
// @Failure			400
// @Router			/book/{id} [put]
func UpdateBook(br *services.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message":     "not found",
				"status_code": "404",
			})
		}

		var b data.Book

		err = c.BindJSON(&b)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":     "bad request",
				"status_code": "400",
			})
		}
		if id != int(b.Id) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":     "bad request",
				"status_code": "400",
			})
		}
		c.JSON(http.StatusOK, br.Update(b.Id, b))
	}
}

// @Summary			Delete Book
//
// @Accept			json
// @Produce			json
// @Param			id	path uint	true "Book ID"
// @Success			200
// @Failure			404
// @Failure			400
// @Failure			500
// @Router			/book/{id} [delete]
func DeleteBook(br *services.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message":     "not found",
				"status_code": "404",
			})
		}
		if br.Delete(uint(id)) {
			c.JSON(http.StatusOK, gin.H{
				"message":     "Ok",
				"status_code": "200",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     "internal error",
			"status_code": "500",
		})
	}
}
