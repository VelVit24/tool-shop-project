package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/VelVit24/projext/dto"
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
		log.Println("here")
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
func (h *ProductHandler) DeleteAdminProducts(c *gin.Context) {
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

func (h *ProductHandler) GetProducts(c *gin.Context) {
	filter := dto.ProductFiler{}
	filter.Page, _ = strconv.Atoi(c.Query("page"))
	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	if value := c.Query("category"); value != "" {
		id, _ := strconv.Atoi(value)
		filter.CategoryID = &id
	}
	if value := c.Query("priceFrom"); value != "" {
		price, _ := strconv.Atoi(value)
		filter.PriceFrom = &price
	}
	if value := c.Query("priceTo"); value != "" {
		price, _ := strconv.Atoi(value)
		filter.PriceTo = &price
	}
	if value := c.Query("inStock"); value != "" {
		stock, _ := strconv.ParseBool(value)
		filter.InStock = &stock
	}
	if value := c.Query("search"); value != "" {
		filter.Search = &value
	}
	if value := c.Query("sort"); value != "" {
		filter.Sort = value
	}
	products, err := h.service.GetProduct(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, gin.H{"products": products, "page": filter.Page, "limit": filter.Limit, "total": len(products)})
}

func (h *ProductHandler) GetProductsId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	product, err := h.service.GetProductId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, product)
}
