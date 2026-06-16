package handlers

import (
	"net/http"
	"strconv"

	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) PostAdminCategories(c *gin.Context) {
	cat := models.Category{}
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.service.CreateCategory(&cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(200)
}

func (h *CategoryHandler) PutAdminCategories(c *gin.Context) {
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
	err = h.service.UpdateCategory(&cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(200)
}
func (h *CategoryHandler) DeleteAdminCategories(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = h.service.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(204)
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	cats, err := h.service.GetCategory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, cats)
}
