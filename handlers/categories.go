package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) POSTCategories(c *gin.Context) {
	cat := models.Category{}
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := repository.InsertCategory(h.DB, &cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(200)
}

func (h *Handler) PUTCategories(c *gin.Context) {
	cat := models.Category{}
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	cat.Id = id
	err = repository.UpdateCategory(h.DB, &cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(200)
}
func (h *Handler) DELETECategories(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = repository.DeleteCategory(h.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(204)
}

func (h *Handler) GETCategories(c *gin.Context) {
	cats, err := repository.SelectCategories(h.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, cats)
}
