package handlers

import (
	"net/http"
	"strconv"

	"github.com/VelVit24/projext/dto"
	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/service"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: *service}
}

func (h *OrderHandler) PostOrders(c *gin.Context) {
	id_user, ok := c.Get("user_id")
	if ok != true {
		c.JSON(http.StatusUnauthorized, "invalid token")
	}
	order := models.Order{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.service.CreateOrder(id_user.(int), &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(200)
}

func (h *OrderHandler) PutAdminOrders(c *gin.Context) {
	order := models.Order{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	order.Id = id
	err = h.service.UpdateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(200)
}
func (h *OrderHandler) DeleteAdminOrders(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = h.service.DeleteOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(204)
}

func (h *OrderHandler) GetAdminOrders(c *gin.Context) {
	role, ok := c.Get("role")
	if ok != true {
		c.JSON(http.StatusUnauthorized, "invalid token")
		return
	}
	page := c.Query("page")
	limit := c.Query("limit")
	orders, err := h.service.GetOrders(-1, role.(string), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, orders)
}

func (h *OrderHandler) GetOrders(c *gin.Context) {
	role, ok := c.Get("role")
	if ok != true {
		c.JSON(http.StatusUnauthorized, "invalid token")
		return
	}
	id_user, ok := c.Get("user_id")
	if ok != true {
		c.JSON(http.StatusUnauthorized, "invalid token")
		return
	}
	page := c.Query("page")
	limit := c.Query("limit")
	orders, err := h.service.GetOrders(id_user.(int), role.(string), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, orders)
}

func (h *OrderHandler) PostOrdersNoAuth(c *gin.Context) {
	order := dto.OrderRequestNoAuth{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.service.CreateOrderOnAuth(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(201)
}
