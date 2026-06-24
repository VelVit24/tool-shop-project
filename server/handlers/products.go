package handlers

import (
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/VelVit24/projext/dto"
	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/service"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "name shoud be unique"})
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(200)
}
func (h *ProductHandler) PostAdminProductSlug(c *gin.Context) {
	var request struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	slug, err := h.service.CreateProductSlug(request.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "slug already exists"})
		return
	}
	c.JSON(200, gin.H{"slug": slug})
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
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "name shoud be unique"})
				return
			}
		}
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
		filter.CategorySlug = &value
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
	c.JSON(200, products)
}

func (h *ProductHandler) GetProductsSlug(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "slug is required"})
		return
	}
	product, err := h.service.GetProductSlug(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, product)
}

func (h *ProductHandler) PostProductImage(c *gin.Context) {
	slug := c.Param("slug")
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get multipart form"})
		return
	}
	files := form.File["images"]
	err = h.service.SetProductImages(slug, &files)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save images"})
		return
	}
	c.Status(200)
}

func (h *ProductHandler) GetProductImage(c *gin.Context) {
	slug := c.Param("slug")
	ind := c.Param("ind")
	size := c.Query("size")
	path := "static/images/products/" + slug
	if size == "small" {
		path += "/small/"
	} else {
		path += "/big/"
	}
	path += ind + ".webp"
	c.File(path)
}
