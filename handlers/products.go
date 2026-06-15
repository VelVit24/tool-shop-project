package handlers

import (
	"net/http"
	"strconv"

	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) PostAdminProduct(c *gin.Context) {
	product := models.Product{}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.service.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(200)
}

func (h *ProductHandler) PutAdminProduct(c *gin.Context) {
	instr := models.Product{}
	if err := c.ShouldBindJSON(&instr); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	instr.Id = id
	err = h.service.UpdateProduct(&instr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(200)
}
func (h *ProductHandler) DeleteAdminInstruments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = h.service.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(204)
}

func (h *ProductHandler) GetAdminInstruments(c *gin.Context) {
	cats, err := h.service.GetProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, cats)
}
