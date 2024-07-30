package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lauro-ss/work-at-olist/internal/services"
)

// @Summary			Author
//
// @Description		list all authors
// @Accept			json
// @Produce			json
// @Success			200	{object}	[]data.Author
// @Router			/author [get]
func ListAuthors(ar *services.AuthorRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, ar.List())
	}
}
