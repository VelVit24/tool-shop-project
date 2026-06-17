package handlers

import (
	"net/http"
	"strconv"

	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/service"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service *service.CartService
}

func NewCartHandler(service *service.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) PostCart(c *gin.Context) {
	id_user, ok := c.Get("user_id")
	if ok != true {
		c.JSON(http.StatusUnauthorized, "invalid token")
	}
	cart := models.Cart{}
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.service.CreateCart(id_user.(int), &cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(200)
}

func (h *CartHandler) PutCart(c *gin.Context) {
	id_user, ok := c.Get("user_id")
	if ok != true {
		c.JSON(http.StatusUnauthorized, "invalid token")
	}
	cart := models.Cart{}
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.service.UpdateCart(id_user.(int), &cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(200)
}
func (h *CartHandler) DeleteCart(c *gin.Context) {
	id_user, ok := c.Get("user_id")
	if ok != true {
		c.JSON(http.StatusUnauthorized, "invalid token")
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = h.service.DeleteCart(id_user.(int), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(204)
}

func (h *CartHandler) GetCart(c *gin.Context) {
	id_user, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, "invalid token")
		return
	}
	carts, err := h.service.GetCart(id_user.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, carts)
}
